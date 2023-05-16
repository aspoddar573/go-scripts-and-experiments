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
	deleteColumnRequest := &tickleDb.UpdateTableRequest{
		Env:       os.Getenv("conf_track"),
		Namespace: os.Getenv("conf_database_name"),
		TableName: "template",
		UpdateRequest: &tickleDb.UpdateTableRequest_DropColumn{
			DropColumn: &tickleDb.DeleteField{
				FieldPath: "thumbnail_media_id",
			},
		},
	}

	deleteTemplateTableColumnResponse, err := helper.UpdateTable(os.Getenv("conf_host")+":"+os.Getenv("conf_port"), deleteColumnRequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	print(deleteTemplateTableColumnResponse)
	waitForSuccessResponse(deleteTemplateTableColumnResponse)
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
