package workflowService

type WorkflowServiceConfig struct {
	WorkflowServiceHost string `envconfig:"WORKFLOW_SERVICE_HOST"`
}
