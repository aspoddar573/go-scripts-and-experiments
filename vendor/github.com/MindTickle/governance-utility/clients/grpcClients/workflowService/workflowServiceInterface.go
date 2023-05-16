package workflowService

import (
	"context"
	"github.com/MindTickle/baton-rabbitmq/protos/generated/go/workflow"
)

type IWorkflowService interface {
	CreateWorkflow(ctx context.Context, orgId int64, jobType string, bulkRequestPayload string) (*workflow.Workflow, error)
}
