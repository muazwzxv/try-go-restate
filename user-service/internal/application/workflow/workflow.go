package workflow

type WorkflowTriggerInfo struct {
	WorkflowTriggerEvent string
	HandlerName          string
	ServiceName          string
	Method               string
	Payload              []byte
}

var (
	WorkflowCreateUser = WorkflowTriggerInfo{
		WorkflowTriggerEvent: "CREATE_USER_WORKFLOW",
		ServiceName:          "UserServiceWorkflows",
		HandlerName:          "ExecuteCreateUserWorkflow",
		Method:               "POST",
	}

	Workflowxxx = WorkflowTriggerInfo{}
)
