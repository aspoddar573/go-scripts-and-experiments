package main

import (
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
	"os"
)

func main() {
	templateCategoryTable := tickleDb.Table{
		TableName: "template_category",
		Ttl:       0,
		Version:   0,
		Namespace: os.Getenv("conf_database_name"),
		Env:       os.Getenv("conf_track"),
		Columns: []*tickleDb.Field{
			&tickleDb.Field{
				FieldName: "name",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      255,
			},
			&tickleDb.Field{
				FieldName: "display_order",
				DataType:  tickleDb.Field_INT,
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

	resp, err := helper.CreateTable(os.Getenv("conf_host")+":"+os.Getenv("conf_port"), &templateCategoryTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
