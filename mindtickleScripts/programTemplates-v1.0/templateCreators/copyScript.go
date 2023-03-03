package templateCreators

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"goProject/mindtickleScripts/programTemplates-v1.0/sqlClient"
	"goProject/mindtickleScripts/programTemplates-v1.0/templateTypes"
	"goProject/mindtickleScripts/utils"
	"strconv"
)

const TEMPLATE_CREATORS_TABLE = "program_template_creator"

func CopyProgramTemplateCreators(targetEnv string, templateCreatorsToCopy []string, additionalInformation map[string]templateTypes.AdditionalTemplateCreatorCopyInformation) {
	fmt.Printf("Request to copy %s template creators to %s environment with %+v information.\n", templateCreatorsToCopy, targetEnv, additionalInformation)
	prodSqlStoreClient := sqlClient.GetTickleDBClient("prod")
	ctx := context.Background()
	response, err := prodSqlStoreClient.Search(ctx, &tickleDbSqlStore.SearchRequest{
		RequestContext: utils.GetRequestContext(1212991592195620381, "prod", "governance"),
		Project:        "governance",
		SqlStatement:   utils.CreateSearchQuery(TEMPLATE_CREATORS_TABLE, templateCreatorsToCopy),
	})
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	rows := response.Rows
	fmt.Printf("%d rows fetched!\n", len(rows))

	var rowsToInsert = make([]*tickleDbSqlStore.Row, 0)
	for _, row := range rows {
		rowDataBytes, err := utils.ConvertRowToDbModel(ctx, row)
		programTemplateCreatorDbRow := &templateTypes.ProgramTemplateCreatorDbRow{}
		err = json.Unmarshal(rowDataBytes, programTemplateCreatorDbRow)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling program template creator: %+v\n", err)
		}
		programTemplateCreatorMetadataRow := &templateTypes.ProgramTemplateCreatorMetadata{}
		err = utils.UnmarshalString(programTemplateCreatorDbRow.Metadata, programTemplateCreatorMetadataRow)
		if err != nil {
			fmt.Printf("Error occurred while marshalling program template creator row: %+v\n", err)
		}
		programTemplateCreatorRow := &templateTypes.ProgramTemplateCreatorRow{
			CreatorId:   programTemplateCreatorDbRow.CreatorId,
			Name:        programTemplateCreatorDbRow.Name,
			Type:        programTemplateCreatorDbRow.Type,
			Description: programTemplateCreatorDbRow.Description,
			Metadata:    *programTemplateCreatorMetadataRow,
		}
		rowValue, err := json.Marshal(programTemplateCreatorRow)
		rowsToInsert = append(rowsToInsert, &tickleDbSqlStore.Row{
			Id: strconv.FormatInt(programTemplateCreatorDbRow.CreatorId, 10),
			RowValue: &tickleDbSqlStore.RowValue{
				RowInBytes: rowValue,
				AuthMeta: &tickleDbSqlStore.AuthMeta{
					GlobalContextId: "golang-script",
					AuthId:          "_mtadmin",
				},
			},
		})
	}
	fmt.Printf("All rows: %+v\n", rowsToInsert)

	targetSqlStoreClient := sqlClient.GetTickleDBClient(targetEnv)
	createRowsResponse, err := targetSqlStoreClient.CreateRows(ctx, &tickleDbSqlStore.CreateRowsRequest{
		RequestContext: utils.GetRequestContext(1212991592195620381, targetEnv, "governance"),
		Project:        "governance",
		TableName:      TEMPLATE_CREATORS_TABLE,
		Rows:           rowsToInsert,
	})
	if err != nil {
		fmt.Printf("Count not create row. Error occurred: %+v\n", err)
	}
	fmt.Printf("Successfully created %d rows\n", createRowsResponse.RowsAffected)
}
