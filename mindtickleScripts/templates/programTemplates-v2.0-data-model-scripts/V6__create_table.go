package main

import (
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
	"os"
)

func main() {
	templateToCategoryMappingTable := tickleDb.Table{
		TableName: "template_to_category_mapping",
		Ttl:       0,
		Version:   0,
		Namespace: os.Getenv("conf_database_name"),
		Env:       os.Getenv("conf_track"),
		Columns: []*tickleDb.Field{
			&tickleDb.Field{
				FieldName: "template_id",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      255,
			},
			&tickleDb.Field{
				FieldName: "category_id",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      255,
			},
		},
		PartitionStrategy: tickleDb.PartitionStrategy_HASH_BASED,
		PartitionKey:      "",
		PrimaryKey: &tickleDb.PrimaryKey{Columns: []string{
			"id",
			"tenant_id",
		}},
		IndexColumns: []*tickleDb.IndexField{
			&tickleDb.IndexField{
				FieldPath: []string{"category_id"},
				IndexName: "category_index",
			},
		},
		UniqueColumns: []*tickleDb.IndexField{
			&tickleDb.IndexField{
				FieldPath: []string{"template_id", "category_id"},
				IndexName: "unique_template_category_index",
			},
		},
	}

	resp, err := helper.CreateTable(os.Getenv("conf_host")+":"+os.Getenv("conf_port"), &templateToCategoryMappingTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
