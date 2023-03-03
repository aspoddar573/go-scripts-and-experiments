package programTemplates

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

const PROGRAM_TEMPLATES_TABLE = "program_template"

func CopyProgramTemplates(targetEnv string, templatesToCopy []string, additionalInformation map[string]templateTypes.AdditionalProgramTemplateCopyInformation) {
	fmt.Printf("Request to copy %s templates to %s environment with %+v information.\n", templatesToCopy, targetEnv, additionalInformation)
	prodSqlStoreClient := sqlClient.GetProdTickleDBClient()
	ctx := context.Background()
	response, err := prodSqlStoreClient.Search(ctx, &tickleDbSqlStore.SearchRequest{
		RequestContext: utils.GetRequestContext(1212991592195620381, "prod", "governance"),
		Project:        "governance",
		SqlStatement:   utils.CreateSearchQuery(PROGRAM_TEMPLATES_TABLE, templatesToCopy),
	})
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	rows := response.Rows
	fmt.Printf("%d rows fetched!\n", len(rows))

	var rowsToInsert = make([]*tickleDbSqlStore.Row, 0)
	for _, row := range rows {
		rowDataBytes, err := utils.ConvertRowToDbModel(ctx, row)
		programTemplateDbRow := &templateTypes.ProgramTemplateDbRow{}
		err = json.Unmarshal(rowDataBytes, programTemplateDbRow)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling program templates: %+v\n", err)
		}

		thumbnail := &templateTypes.Thumbnail{}
		listingThumbnail := &templateTypes.Thumbnail{}
		var competencies []*templateTypes.CompetencyDBDocument
		var frequentlyAskedQuestions []*templateTypes.FrequentlyAskedQuestion
		var setUpGuidelines []*templateTypes.SetupGuideline
		var templateSupportingDocuments []*templateTypes.TemplateSupportingDocument

		err = utils.UnmarshalString(programTemplateDbRow.Thumbnail, thumbnail)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling thumbnail object: %+v\n", err)
		}
		err = utils.UnmarshalString(programTemplateDbRow.ListingThumbnail, listingThumbnail)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling listing thumbnail object: %+v\n", err)
		}
		err = utils.UnmarshalString(programTemplateDbRow.Competencies, &competencies)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling competencies object: %+v\n", err)
		}
		err = utils.UnmarshalString(programTemplateDbRow.FrequentlyAskedQuestions, &frequentlyAskedQuestions)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling frequently asked questions object: %+v\n", err)
		}
		err = utils.UnmarshalString(programTemplateDbRow.SetupGuidelines, &setUpGuidelines)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling setup guidelines object: %+v\n", err)
		}
		err = utils.UnmarshalString(programTemplateDbRow.TemplateSupportingDocuments, &templateSupportingDocuments)
		if err != nil {
			fmt.Printf("Error occurred while unmarshalling template supporting documents object: %+v\n", err)
		}

		programTemplateRow := &templateTypes.ProgramTemplateRow{
			TemplateId:                               programTemplateDbRow.TemplateId,
			CompanyId:                                utils.GetCompanyIdForEnv(targetEnv),
			TemplateSeriesId:                         additionalInformation[strconv.FormatInt(programTemplateDbRow.TemplateId, 10)].SeriesId,
			Name:                                     programTemplateDbRow.Name,
			Purpose:                                  programTemplateDbRow.Purpose,
			Description:                              programTemplateDbRow.Description,
			PostCreationTitle:                        programTemplateDbRow.PostCreationTitle,
			PostCreationDescription:                  programTemplateDbRow.PostCreationDescription,
			State:                                    programTemplateDbRow.State,
			CreatedBy:                                programTemplateDbRow.CreatedBy,
			UpdatedBy:                                programTemplateDbRow.UpdatedBy,
			EstimatedCompletionTimeInSeconds:         programTemplateDbRow.EstimatedCompletionTimeInSeconds,
			EstimatedCompletionTimeIntervalInSeconds: programTemplateDbRow.EstimatedCompletionTimeIntervalInSeconds,
			Scope:                                    programTemplateDbRow.Scope,
			CreatorId:                                programTemplateDbRow.CreatorId,
			Thumbnail:                                thumbnail,
			ListingThumbnail:                         listingThumbnail,
			Competencies:                             competencies,
			SetupGuidelines:                          setUpGuidelines,
			FrequentlyAskedQuestions:                 frequentlyAskedQuestions,
			TemplateSupportingDocuments:              templateSupportingDocuments,
		}

		rowValue, err := json.Marshal(programTemplateRow)
		if err != nil {
			fmt.Printf("Error occurred while marshalling program templates: %+v\n", err)
		}
		rowsToInsert = append(rowsToInsert, &tickleDbSqlStore.Row{
			Id: strconv.FormatInt(programTemplateRow.TemplateId, 10),
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
		TableName:      PROGRAM_TEMPLATES_TABLE,
		Rows:           rowsToInsert,
	})
	if err != nil {
		fmt.Printf("Count not create row. Error occurred: %+v\n", err)
	}
	fmt.Printf("Successfully created %d rows\n", createRowsResponse.RowsAffected)
}
