package main

import (
	"context"
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
	"os"
	"time"
)

const ResponseStatusCheckInterval = 10 * time.Second
const MaxRetryCount = 180

func main() {
	updateTemplateTableRequest := &tickleDb.UpdateTableRequest{
		Env:       os.Getenv("conf_track"),
		Namespace: os.Getenv("conf_database_name"),
		TableName: "template",
		UpdateRequest: &tickleDb.UpdateTableRequest_AddColumn{
			AddColumn: &tickleDb.AddField{
				Field: &tickleDb.Field{
					FieldName: "template_metadata",
					DataType:  tickleDb.Field_JSON,
					Required:  false,
					NestedFields: []*tickleDb.Field{ // List of Objects
						{
							FieldName: "key",
							DataType:  tickleDb.Field_STRING,
							Size:      255,
							Required:  false,
						},
					},
				},
			},
		},
	}

	updateTemplateTableResponse, err := helper.UpdateTable(os.Getenv("conf_host")+":"+os.Getenv("conf_port"), updateTemplateTableRequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	print(updateTemplateTableResponse)
	waitForSuccessResponse(updateTemplateTableResponse)
}

func waitForSuccessResponse(resp *tickleDb.UpdateTableResponse) {
	client, err := helper.GetClient(os.Getenv("conf_host") + ":" + os.Getenv("conf_port"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	getUpdateTableStatusRequest := tickleDb.GetUpdateTableStatusRequest{
		UpdateTableTaskId: resp.UpdateTableTaskId,
	}
	for i := 0; i < MaxRetryCount; i++ {
		getUpdateTableStatusResponse, err := client.GetUpdateTableStatus(context.Background(), &getUpdateTableStatusRequest)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if getUpdateTableStatusResponse.Status == "FAILED" {
			fmt.Printf("Got failed status for %+v. Exiting the update table flow.", getUpdateTableStatusRequest)
			os.Exit(1)
		}
		if getUpdateTableStatusResponse.Status == "COMPLETED" {
			return
		}
		time.Sleep(ResponseStatusCheckInterval)
	}
	fmt.Printf("Couldn't find status for %+v in %d retry attempts. Exiting the update table flow.", getUpdateTableStatusRequest, MaxRetryCount)
	os.Exit(1)
}
