package main

import (
	"fmt"
	"os"

	"github.com/MindTickle/storageprotos/pb/tickleDb"
	helper "github.com/MindTickle/tickledb-data-automation"
)

func main() {
	templateTable := tickleDb.Table{
		TableName: "template",
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
				FieldName:          "type",
				DataType:           tickleDb.Field_ENUMERATION,
				Required:           true,
				EnumExpectedValues: []string{"PROGRAM", "MODULE"},
				DefaultValue:       "PROGRAM",
			},
			&tickleDb.Field{
				FieldName:          "scope",
				DataType:           tickleDb.Field_ENUMERATION,
				Required:           true,
				EnumExpectedValues: []string{"OPEN", "RESTRICTED", "CLOSED"},
				DefaultValue:       "PROGRAM",
			},
			&tickleDb.Field{
				FieldName:          "state",
				DataType:           tickleDb.Field_ENUMERATION,
				Required:           true,
				EnumExpectedValues: []string{"DRAFT", "PUBLISHED", "ARCHIVED"},
				DefaultValue:       "DRAFT",
			},
			&tickleDb.Field{
				FieldName: "creator_id",
				DataType:  tickleDb.Field_STRING,
				Required:  true,
				Size:      255,
			},
			&tickleDb.Field{
				FieldName: "flags",
				DataType:  tickleDb.Field_INT,
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "purpose",
				DataType:  tickleDb.Field_STRING,
				Required:  false,
				Size:      2000,
			},
			&tickleDb.Field{
				FieldName: "description",
				DataType:  tickleDb.Field_STRING,
				Required:  false,
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
				FieldName: "post_creation_intro_title",
				DataType:  tickleDb.Field_STRING,
				Required:  false,
				Size:      255,
			},
			&tickleDb.Field{
				FieldName: "post_creation_intro_description",
				DataType:  tickleDb.Field_STRING,
				Required:  false,
				Size:      2000,
			},
			&tickleDb.Field{
				FieldName: "preview",
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
			&tickleDb.Field{
				FieldName: "setup_guidelines",
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
			&tickleDb.Field{
				FieldName: "frequently_asked_questions",
				DataType:  tickleDb.Field_JSON,
				Required:  false,
				NestedFields: []*tickleDb.Field{ // List of Objects
					{
						FieldName: "key",
						DataType:  tickleDb.Field_STRING,
						Size:      1000,
						Required:  false,
					},
				},
			},
			&tickleDb.Field{
				FieldName: "view_count",
				DataType:  tickleDb.Field_INT64,
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "company_id",
				DataType:  tickleDb.Field_INT64, // can have _default value
				Required:  true,
			},
			&tickleDb.Field{
				FieldName: "template_specific_info",
				DataType:  tickleDb.Field_JSON,
				Required:  false,
				NestedFields: []*tickleDb.Field{ // List of Objects
					{
						FieldName: "key",
						DataType:  tickleDb.Field_STRING,
						Size:      1000,
						Required:  false,
					},
				},
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
				FieldPath: []string{"scope"},
				IndexName: "scope_index",
			},
			&tickleDb.IndexField{
				FieldPath: []string{"state", "type"},
				IndexName: "state_type_index",
			},
		},
		UniqueColumns: []*tickleDb.IndexField{
			&tickleDb.IndexField{
				FieldPath: []string{"creator_id", "type", "name"},
				IndexName: "unique_creator_type_name_index",
			},
		},
	}
	resp, err := helper.CreateTable((os.Getenv("conf_host") + ":" + os.Getenv("conf_port")), &templateTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
