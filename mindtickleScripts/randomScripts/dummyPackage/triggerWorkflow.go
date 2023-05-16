package dummyPackage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MindTickle/governance-protos/pb/common"
	"github.com/MindTickle/governance-protos/pb/templateWorkflows"
	"github.com/MindTickle/governance-utility/clients/grpcClients/workflowService"
	"github.com/MindTickle/governance-utility/govConstants"
	"github.com/MindTickle/governance-utility/logger"
)

func NewMockContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, govConstants.COMPANY_ID, "_testCompanyId")
	ctx = context.WithValue(ctx, govConstants.ORG_ID, "12345")
	ctx = context.WithValue(ctx, govConstants.AUTHORIZER_ID, "_testUserId")
	ctx = context.WithValue(ctx, govConstants.CORRELATION_ID, "_testCallerReqId")
	ctx = context.WithValue(ctx, govConstants.CALLER_REQ_ID, "_testCallerReqId")
	ctx = context.WithValue(ctx, govConstants.APPNAME, "_testCalleeReqId")
	return ctx
}

func TriggerAWorkflow() {
	loggerConfig := logger.Configuration{
		LogLevel:   logger.DEBUG,
		EnableJSON: true,
		Output:     logger.CONSOLE,
	}
	err := logger.NewLogger(loggerConfig, logger.InstanceZapLogger)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := NewMockContext()
	workflowServiceClient, err := workflowService.GetWorkflowServiceClient(ctx, &workflowService.WorkflowServiceConfig{
		WorkflowServiceHost: "btn-svc-workflow.internal-grpc.integration.mindtickle.com:80",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	useModuleTemplateReq := templateWorkflows.CreateModuleFromTemplate{
		TemplateId: "1680681673659707000",
		ModuleName: "module from template in milestone take 4",
		Series: &templateWorkflows.Series{
			IsNew: false,
			Id:    "1648282048091060438",
			//IsNew:      true,
			//StringName: "naya banaya hai series",
		},
		Meta: &common.Meta{
			CompanyId:     "1574261193079891874",
			OrgId:         "1574261180771651819",
			AuthorizerId:  "15d8e6383670c000",
			CorrelationId: "script-test",
		},
	}
	bytes, err := json.Marshal(useModuleTemplateReq)
	fmt.Println("request constructed")
	workflow, err := workflowServiceClient.CreateWorkflow(ctx, 1587814705391751830, string(govConstants.CREATE_MODULE_FROM_TEMPLATE), string(bytes))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("workflow successful")
	fmt.Println(workflow)
}
