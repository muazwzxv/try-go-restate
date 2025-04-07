# Learning Restate with Go 

[https://docs.restate.dev](https://docs.restate.dev)


## How I envision a backend application with restate server would look like

```
┌─────────────────────────────────────────────────┐
│                 Go Service                      │
├─────────────────┬───────────────────────────────┤
│  Experience API │   Workflow API (Restate)      │
│  (Fast Path)    │   (Slow Path)                 │
└─────────────────┴───────────────────────────────┘
        │                         │
        │                         │
        ▼                         ▼
┌───────────────┐       ┌─────────────────────┐
│  DB/Cache     │       │ Restate Server      │
└───────────────┘       └─────────────────────┘
```

### App running in a single backend deployment
- An API server to serve all our APIs (experience layer & workflow API)
- We can use any Mux preferred (Gin, Chi, GoFr, echo)
- Restate server (register all our restate handlers to serve long running workflows)
- Potentially consumers if we have any queues or kafka streams in our application

### Experience layer API
- API's that are used to serve screens (profile page, listings, details page)
- API expected to be simple reads/aggregation of data from our DB and some caching mechanism to reduce loads on DB

### Workflow layer API
- Workflow layers API are treated as mechanism for client to trigger any long running workflows
- Upon API hit, we'll offload the long running running task to restate, the logic layer for worfklow layer API is to do validations necessary prior to triggering restate handlers

