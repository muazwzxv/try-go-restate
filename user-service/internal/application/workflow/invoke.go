package workflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type Invoker interface {
	InvokeWorkflow(ctx context.Context, params *InvocationParams) (*InvocationResponse, error)
}

type InvocationHandler struct {
	httpClient *http.Client
	url        string
}

type InvocationParams struct {
	Workflow WorkflowTriggerInfo
	Payload  []byte
	// use a stable uuid as an idempotency key; Restate deduplicates for us
	IdempotencyKey string
}

type ClientConfig struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

// NewInvocationHandler creates a new handler with a configured HTTP client
func NewInvocationHandler(baseURL string, config *ClientConfig) *InvocationHandler {
	// Use default config if none provided
	cfg := DefaultConfig()
	if config != nil {
		cfg = *config
	}

	// Create a transport with optimized connection pooling
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          cfg.MaxIdleConns,
		MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
		IdleConnTimeout:       cfg.IdleConnTimeout,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   cfg.Timeout,
	}

	return &InvocationHandler{
		httpClient: client,
		url:        baseURL,
	}
}

func DefaultConfig() ClientConfig {
	return ClientConfig{
		Timeout:             30 * time.Second,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}
}

var RESTATE_URL = ""

// Durable RPC call to the product service
func (i *InvocationHandler) InvokeWorkflow(ctx context.Context, params *InvocationParams) (*InvocationResponse, error) {
	// Restate registers the request and makes sure it runs to completion exactly once
	// This is a call to Virtual Object so we can be sure only one reservation is made concurrently
	var (
		httpReq *http.Request
		err     error
	)
	url := fmt.Sprintf("%s/%s/%s", RESTATE_URL, params.Workflow.ServiceName, params.Workflow.HandlerName)
	if params.Payload != nil {
		httpReq, err = http.NewRequest(params.Workflow.Method, url, nil)
		if err != nil {
			slog.Error("Failed creating http request", "err", err.Error())
			return nil, err
		}
	} else {
		httpReq, err = http.NewRequest(params.Workflow.Method, url, bytes.NewBuffer(params.Payload))
		if err != nil {
			slog.Error("Failed creating http request", "err", err.Error())
			return nil, err
		}
	}

	if params.IdempotencyKey != "" {
		httpReq.Header.Set("idempotency-key", params.IdempotencyKey)
	}

	resp, err := i.httpClient.Do(httpReq)
	if err != nil {
		slog.Error("Failed calling http request", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed parsing http response", "err", err.Error())
		return nil, err
	}

	var invocationResponse InvocationResponse
	err = json.Unmarshal(body, &invocationResponse)
	if err != nil {
		slog.Error("Failed unmarshaling response into struct", "err", err.Error())
		return nil, err
	}

	slog.Info("Parsed invocation",
		"invocationId", invocationResponse.InvocationID,
		"status", invocationResponse.Status)

	return &invocationResponse, nil
}

type InvocationResponse struct {
	InvocationID string `json:"invocationId"`
	Status       string `json:"status"`
}
