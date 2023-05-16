package workflowService

import (
	"github.com/MindTickle/baton-rabbitmq/protos/generated/go/workflow"
	govClientHelper "github.com/MindTickle/governance-utility/clients/helper"
	"github.com/MindTickle/governance-utility/govConstants"
	"github.com/MindTickle/governance-utility/helper"
	"github.com/MindTickle/governance-utility/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkflowServiceImpl struct {
	client      workflow.WorkflowServiceClient
	workflowEnv string
}

func GetWorkflowServiceClient(ctx context.Context, conf *WorkflowServiceConfig) (IWorkflowService, error) {
	workflowConn, err := grpc.Dial(conf.WorkflowServiceHost, grpc.WithInsecure())
	if err != nil {
		logger.Error(ctx, logger.NewFacets(), "error connectiong workflowConn :%s", err)
		return nil, status.Errorf(codes.Unavailable, "workflowConn connection not established")
	}

	return &WorkflowServiceImpl{
		client:      workflow.NewWorkflowServiceClient(workflowConn),
		workflowEnv: govConstants.ENV_VAR_TRACK,
	}, nil
}

func (w *WorkflowServiceImpl) CreateWorkflow(ctx context.Context, orgId int64, jobType string, bulkRequestPayload string) (*workflow.Workflow, error) {
	logger.Infof(ctx, logger.NewFacets(), "Response from workflow service is", w.client, " => ", w.workflowEnv)

	ctxWithMeta := govClientHelper.InjectSpanIntoContextForRpc(ctx)
	response, err := w.client.CreateWorkflow(ctxWithMeta, &workflow.CreateWorkflowRequest{
		Workflow: &workflow.Workflow{
			Namespace:       govConstants.PROJECT,
			OrgId:           helper.OrgToStringWithoutError(orgId),
			AuthId:          "governance_workflow",
			BulkRequestType: jobType,
			BulkRequest:     bulkRequestPayload,
			WorkflowState:   workflow.WorkflowState_WORKFLOW_CREATED,
		},
	})
	if err != nil {
		logger.Error(ctx, logger.NewFacets(), "error from workflow service ", err)
		return nil, err
	}

	logger.Infof(ctx, logger.NewFacets(), "Response from workflow service is", response)
	return response.Workflow, err
}
