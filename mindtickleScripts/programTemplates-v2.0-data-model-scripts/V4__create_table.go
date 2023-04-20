package main

import (
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
	"os"
)

func main() {
	templateAccessibilityTable := tickleDb.Table{
		TableName: "template_accessibility",
		Ttl:       0,
		Version:   0,
		Namespace: os.Getenv("conf_database_name"),
		Env:       os.Getenv("conf_track"),
		Columns: []*tickleDb.Field{
			&tickleDb.Field{
				FieldName: "company_id",
				DataType:  tickleDb.Field_INT64,
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "template_id",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      255,
			},
			&tickleDb.Field{
				FieldName: "is_visible", /* to control visibility of program templates */
				DataType:  tickleDb.Field_BOOL,
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "is_usable", /* to control usability of program templates */
				DataType:  tickleDb.Field_BOOL,
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
				FieldPath: []string{"company_id", "template_id"},
				IndexName: "company_template_index",
			},
		},
	}

	resp, err := helper.CreateTable(os.Getenv("conf_host")+":"+os.Getenv("conf_port"), &templateAccessibilityTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
