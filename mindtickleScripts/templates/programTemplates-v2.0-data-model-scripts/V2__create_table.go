package main

import (
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
	"os"
)

func main() {
	templateCreatorTable := tickleDb.Table{
		TableName: "template_creator",
		Ttl:       0,
		Version:   0,
		Namespace: "templates",
		Env:       "integration",
		Columns: []*tickleDb.Field{
			&tickleDb.Field{
				FieldName: "name",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      255,
			},
			&tickleDb.Field{
				FieldName:          "type",
				DataType:           tickleDb.Field_ENUMERATION,
				Required:           true,
				EnumExpectedValues: []string{"MINDTICKLE", "PARTNER", "CUSTOMER"},
				DefaultValue:       "MINDTICKLE",
			},
			&tickleDb.Field{
				FieldName: "description",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      2000,
			},
			&tickleDb.Field{
				FieldName: "listing_thumbnail_media_id",
				DataType:  tickleDb.Field_INT64,
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "thumbnail_media_id",
				DataType:  tickleDb.Field_INT64,
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "metadata",
				DataType:  tickleDb.Field_JSON,
				Required:  false,
				NestedFields: []*tickleDb.Field{
					{
						FieldName: "key",
						DataType:  tickleDb.Field_STRING,
						Size:      255,
						Required:  false,
					},
				},
			},
			&tickleDb.Field{
				FieldName: "company_id",
				DataType:  tickleDb.Field_INT64,
				Required:  true,
			},
		},
		PartitionStrategy: tickleDb.PartitionStrategy_HASH_BASED,
		PartitionKey:      "",
		PrimaryKey: &tickleDb.PrimaryKey{Columns: []string{
			"id",
			"tenant_id",
		}},
		UniqueColumns: []*tickleDb.IndexField{
			&tickleDb.IndexField{
				FieldPath: []string{"name"},
				IndexName: "unique_name_index",
			},
		},
	}

	resp, err := helper.CreateTable("tdb-svc-manager.internal-grpc.integration.mindtickle.com"+":"+"80", &templateCreatorTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
