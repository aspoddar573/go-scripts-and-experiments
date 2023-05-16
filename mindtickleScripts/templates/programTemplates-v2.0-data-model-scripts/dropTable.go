package main

import (
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
)

func main() {
	resp, err := helper.DeleteTable("tdb-svc-manager.internal-grpc.integration.mindtickle.com"+":"+"80", &tickleDb.DeleteTableRequest{
		Env:       "integration",
		Namespace: "templates",
		TableName: "template_creator",
		Tenant:    &tickleDb.DeleteTableRequest_TAll{},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
