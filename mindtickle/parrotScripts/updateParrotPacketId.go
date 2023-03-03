package parrotScripts

import (
	"context"
	"fmt"
	"github.com/MindTickle/governance-protos/pb/feedback"
	"github.com/MindTickle/governance-protos/pb/governanceFeedback"
	"github.com/MindTickle/governance-protos/pb/parrot"
	"github.com/MindTickle/governance-utility/govConstants"
	"github.com/MindTickle/governance-utility/logger"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sync"
)

const PARROT_HOST = "gvn-svc-parrot.internal-grpc.prod.mindtickle.com:80"
const BULK_OP_APPNAME = "bulk-ops-baton-workflow"

type Resource string

const (
	BULK_COPY_MODULES Resource = "BULK_COPY_MODULES"
)

var parrotServiceOnce sync.Once
var parrotServiceConn *grpc.ClientConn

func GetParrotServiceClient(ctx context.Context) (parrot.ParrotServiceClient, error) {
	var err error
	parrotServiceOnce.Do(func() {
		fmt.Println("Intialising Parrot Service client, %s", PARROT_HOST)
		serviceHost := PARROT_HOST
		parrotServiceConn, err = grpc.Dial(serviceHost, grpc.WithInsecure())

	})
	if err != nil {
		fmt.Println("error establishing connection with parrot service : %s", err)
		return nil, err
	}

	return parrot.NewParrotServiceClient(parrotServiceConn), nil
}

func GetBulkOpWorkflowStatus(numOps int, numSuccess int, numFailed int) feedback.Status {
	var status feedback.Status
	if numSuccess+numFailed < numOps {
		status = feedback.Status_IN_PROGRESS
	} else if numFailed == 0 {
		status = feedback.Status_SUCCESS
	} else if numSuccess == 0 {
		status = feedback.Status_FAILED
	} else {
		status = feedback.Status_PARTIALLY_FAILED
	}
	return status
}

func GetBulkOpFeedbackPayload(ctx context.Context, status feedback.Status, totalOps int, successOps int, failedOps int, feedbackDetails *governanceFeedback.BulkModuleOpsFeedbackDetails) (*any.Any, error) {
	progress := &feedback.Progress{
		IsEnabled: true,
		Total:     int32(totalOps),
		Completed: int32(successOps),
		Failed:    int32(failedOps),
	}

	feedbackDetailsMarshalled, err := ptypes.MarshalAny(feedbackDetails)
	if err != nil {
		logger.Errorf(ctx, logger.NewFacets(), "Error in details marshal, error = %+v", err)
		return nil, err
	}

	details := &feedback.Details{
		IsEnabled:       true,
		DetailsExpanded: feedbackDetailsMarshalled,
	}

	remainingOps := totalOps - (successOps + failedOps)
	estimatedTime := int32(remainingOps * 5 / 2) // 2.5 sec for each operation
	feedbackObj := &feedback.Feedback{
		Status:                         status,
		Progress:                       progress,
		Details:                        details,
		EstimatedCompletionTimeSeconds: estimatedTime,
	}
	payload, err := ptypes.MarshalAny(feedbackObj)
	if err != nil {
		logger.Errorf(ctx, logger.NewFacets(), "Error in json marshal, error = %+v", err)
		return nil, err
	}
	return payload, nil
}

func GetContextForParrotRequest(ctx context.Context) context.Context {
	ctxWithMeta := metadata.AppendToOutgoingContext(ctx,
		govConstants.COMPANY_ID, "1610305323933878854",
		govConstants.AUTHORIZER_ID, "160abaf3a351e000",
		govConstants.ORG_ID, "1584698402418200979",
		govConstants.CORRELATION_ID, "manual-gepi-1960",
		govConstants.APPNAME, "governance-processors",
	)
	return ctxWithMeta
}

func GetBulkOpTemplateKey(resource Resource, status feedback.Status, singleOp bool) *parrot.MessageTemplateKey {
	templateName := fmt.Sprintf("%s-%s", resource, status.String())
	if singleOp {
		templateName = fmt.Sprintf("%s-SINGLE", templateName)
	}
	var version int32 = 1
	// Add if condition here to use different versions for different templates
	return &parrot.MessageTemplateKey{
		Namespace:       "governance",
		ApplicationName: BULK_OP_APPNAME,
		TemplateName:    templateName,
		Version:         version,
	}
}

func GetCopyTemplateResolvers(totalModules int, failedModules int, totalSeries int) []*parrot.MessageTemplateResolver {
	resolvers := []*parrot.MessageTemplateResolver{
		{
			Key:   "totalModules",
			Value: fmt.Sprintf("%d", totalModules),
		},
		{
			Key:   "failedModules",
			Value: fmt.Sprintf("%d", failedModules),
		},
		{
			Key:   "totalSeries",
			Value: fmt.Sprintf("%d", totalSeries),
		},
	}
	return resolvers
}

func GetCopyOpsFeedback(moduleId string, sourceSeriesId string, destSeriesId string, isNewSeries bool, seriesName string, status feedback.Status, errorCode string) *governanceFeedback.ModuleOpsFeedback {
	copyFeedback := &governanceFeedback.ModuleCopyFeedback{
		ModuleId:            moduleId,
		SourceSeriesId:      sourceSeriesId,
		DestinationSeriesId: destSeriesId,
		NewSeries:           isNewSeries,
		SeriesName:          seriesName,
		Status:              status,
		FailureReason:       errorCode,
	}
	return &governanceFeedback.ModuleOpsFeedback{
		ModuleOpsFeedback: &governanceFeedback.ModuleOpsFeedback_ModuleCopyFeedback{ModuleCopyFeedback: copyFeedback},
	}
}

func UpdatePacket() error {
	var packetId int64 = 1618770265887979968

	ctx := context.Background()
	parrotClient, err := GetParrotServiceClient(ctx)
	if err != nil {
		fmt.Println("error occurred!")
		return err
	}

	moduleOpsFeedback := make([]*governanceFeedback.ModuleOpsFeedback, 0)

	moduleOpsFeedback = append(moduleOpsFeedback,
		GetCopyOpsFeedback("1618705671279503688", "1612860449083319315", "1612860449083319315", false, "", feedback.Status_FAILED, ""))

	feedbackDetails := &governanceFeedback.BulkModuleOpsFeedbackDetails{
		ModuleOpsFeedback: moduleOpsFeedback,
	}

	fmt.Println("Sending Request for updating packet %d", packetId)
	workflowStatus := GetBulkOpWorkflowStatus(1, 0, 1)
	payload, err := GetBulkOpFeedbackPayload(ctx, workflowStatus, 1, 0, 1, feedbackDetails)
	ctxWithMeta := GetContextForParrotRequest(ctx)
	fmt.Println("Sent Request for updating packet %d", packetId)
	_, err = parrotClient.UpdatePacket(ctxWithMeta, &parrot.UpdatePacketRequest{
		PacketId:                 packetId,
		Payload:                  payload,
		MessageTemplateKey:       GetBulkOpTemplateKey(BULK_COPY_MODULES, workflowStatus, true),
		MessageTemplateResolvers: GetCopyTemplateResolvers(1, 1, 1),
	})
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Successfully Updated!")
	}
	return nil
}

func FetchPacket() {
	var packetId int64 = 1618770265887979968

	ctx := context.Background()
	parrotClient, err := GetParrotServiceClient(ctx)
	if err != nil {
		fmt.Println("error occurred!")
	}
	ctxWithMeta := GetContextForParrotRequest(ctx)
	fmt.Println("Fetching packet %d", packetId)
	packets, err := parrotClient.BulkFetchPacket(ctxWithMeta, &parrot.BulkFetchPacketRequest{
		PacketIds: []int64{packetId},
	})
	packet := packets.Packets[0]
	feedbackPacket := new(feedback.Feedback)
	_ = ptypes.UnmarshalAny(packet.Payload, feedbackPacket)
	bulkModuleFeedback := new(governanceFeedback.BulkModuleOpsFeedbackDetails)
	_ = ptypes.UnmarshalAny(feedbackPacket.Details.DetailsExpanded, bulkModuleFeedback)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bulkModuleFeedback)
	}

}
