package main

import (
	"fmt"
	"github.com/MindTickle/governance-protos/pb/common"
	"github.com/MindTickle/governance-protos/pb/healthCheckService"
	"github.com/MindTickle/governance-protos/pb/programs"
	"github.com/MindTickle/governance-utility/govConstants"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	ctx := context.TODO()

	var requiredFields = []string{
		govConstants.ORG_ID,
		govConstants.COMPANY_ID,
		govConstants.CORRELATION_ID,
		govConstants.AUTHORIZER_ID,
	}

	for _, reqField := range requiredFields {
		ctx = metadata.AppendToOutgoingContext(ctx, reqField, "1")
	}

	ctx = metadata.AppendToOutgoingContext(ctx, govConstants.APPNAME, "self-health-check")

	service := "program-service.internal-grpc.prod.mindtickle.com"
	port := "443"
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", service, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUserAgent("AGENT-SMITH"))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := healthCheckService.NewHealthCheckClient(conn)
	cu := common.EmptyMsg{}

	response, err := c.HealthCheck(ctx, &cu)
	if err != nil {
		log.Fatalf("health check failed: %s", err)
	}

	log.Printf("Response from server: %s", response)

	if response.Healthy != true {
		panic(nil)
	}

	pg := programs.NewProgramServiceClient(conn)
	updateResponse, err := pg.UpdateProgramTemplateCreator(ctx, &programs.UpdateProgramTemplateCreatorRequest{
		CreatorId:   "1650262556293827452",
		Name:        "Mindtickle CaaS",
		Description: "Mindtickle, a leader in the sales readiness space helps fast-growing companies prepare their sales teams and partners in a scalable and effective way. Mindtickle's Content as a Service (CaaS) team develops impactful enablement programs leveraging Mindtickle, to boost the engagement and adoption of enablement content. CaaS has built 500+ programs for leading organizations from multiple industries such as Pharma, Finance, Technology, FMCG etc.",
		Metadata: &programs.ProgramTemplateCreatorMetadata{
			Email:   "caas@mindtickle.com",
			Website: "mindtickle.com",
			CompanyLogo: &programs.Thumbnail{
				MediaType:            programs.ThumbnailMediaType_IMAGE,
				OriginalUrl:          "https://assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/original.png",
				ProcessedUrl:         "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/processed.png",
				ProcessedUrl_180X120: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/processed_180_120.png",
				ProcessedUrl_600X360: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/processed_600_360.png",
			},
			CompanyListingLogo: &programs.Thumbnail{
				MediaType:            programs.ThumbnailMediaType_IMAGE,
				OriginalUrl:          "https://assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/original.jpeg",
				ProcessedUrl:         "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/processed.jpeg",
				ProcessedUrl_180X120: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/processed_180_120.jpeg",
				ProcessedUrl_600X360: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/processed_600_360.jpeg",
			},
		},
	})
	fmt.Printf("The response is %s\n", updateResponse.OperationSuccessful)

	//creatorResponse, err := pg.CreateProgramTemplateCreator(ctx, &programs.CreateProgramTemplateCreatorRequest{
	//	Name:        "Mindtickle CaaS",
	//	Type:        programs.ProgramTemplateCreatorType_MINDTICKLE,
	//	Description: "Mindtickle, a leader in the sales readiness space helps fast-growing companies prepare their sales teams and partners in a scalable and effective way. Mindtickle's Content as a Service (CaaS) team develops impactful enablement programs leveraging Mindtickle, to boost the engagement and adoption of enablement content. CaaS has built 500+ programs for leading organizations from multiple industries such as Pharma, Finance, Technology, FMCG etc.",
	//	Metadata: &programs.ProgramTemplateCreatorMetadata{
	//		Email:   "caas@mindtickle.com",
	//		Website: "www.mindtickle.com",
	//		CompanyLogo: &programs.Thumbnail{
	//			MediaType:            programs.ThumbnailMediaType_IMAGE,
	//			OriginalUrl:          "https://assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/original.png",
	//			ProcessedUrl:         "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/processed.png",
	//			ProcessedUrl_180X120: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/processed_180_120.png",
	//			ProcessedUrl_600X360: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/main_logo/processed_600_360.png",
	//		},
	//		CompanyListingLogo: &programs.Thumbnail{
	//			MediaType:            programs.ThumbnailMediaType_IMAGE,
	//			OriginalUrl:          "https://assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/original.jpeg",
	//			ProcessedUrl:         "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/processed.jpeg",
	//			ProcessedUrl_180X120: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/processed_180_120.jpeg",
	//			ProcessedUrl_600X360: "//assets.mindtickle.com/program-template-assets/MindTickle_Caas_Creator/listing_logo/processed_600_360.jpeg",
	//		},
	//	},
	//})
	//fmt.Printf("The creator id is %s\n", creatorResponse.CreatorId)

}
