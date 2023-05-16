package main

import "goScriptsAndExperiments/mindtickleScripts/randomScripts/updateTableWithNewMedia"

// 1679553590166624000
// 1679554169649215000

type ProgramTemplateCreator struct {
	CreatorId int64 `json:"creator_id,omitempty,int64"`
}

type TemplateCategory struct {
	DisplayOrder int64 `json:"display_order,omitempty,int"`
}

func main() {
	//ctx := context.Background()
	//targetEnv := "integration"
	//targetSqlStoreClient := sqlClient.GetTickleDBClient(targetEnv)

	//searchRes, err := targetSqlStoreClient.Search(ctx, &tickleDbSqlStore.SearchRequest{
	//	RequestContext: utils.GetRequestContext(1, targetEnv, "templates"),
	//	SqlStatement:   "select * from template_creator",
	//})
	//fmt.Println(searchRes)

	// Default template org id: 1583043504811348643
	// Dummy template org id: 1583043504811348643
	//deleteRes, err := targetSqlStoreClient.DeleteRowById(ctx, &tickleDbSqlStore.DeleteRowByIdRequest{
	//	RequestContext: utils.GetRequestContext(1583043504811348643, targetEnv, "templates"),
	//	Project:        "templates",
	//	TableName:      "template",
	//	Id:             "1679755956986730000",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(deleteRes)

	//category := TemplateCategory{DisplayOrder: 1000000}
	//rowValue, err := json.Marshal(category)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//updateRes, err := targetSqlStoreClient.UpdateRowById(ctx, &tickleDbSqlStore.UpdateRowByIdRequest{
	//	RequestContext: utils.GetRequestContext(1, targetEnv, "templates"),
	//	Project:        "templates",
	//	TableName:      "template_category",
	//	Row: &tickleDbSqlStore.Row{
	//		Id: "1679919939480251000",
	//		RowValue: &tickleDbSqlStore.RowValue{
	//			RowInBytes: rowValue,
	//			AuthMeta: &tickleDbSqlStore.AuthMeta{
	//				GlobalContextId: "aspoddar-local",
	//				AuthId:          "_mt_admin",
	//			},
	//		},
	//	},
	//	Mask: &tickleDbSqlStore.Mask{Fields: []string{"display_order"}},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(updateRes)

	//execRes, err := targetSqlStoreClient.Exec(ctx, &tickleDbSqlStore.ExecRequest{
	//	RequestContext: utils.GetRequestContext(1, targetEnv, "templates"),
	//	SqlStatements:  []string{"DELETE FROM template_creator as tc WHERE tc.id='1679553590166624000'"},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(execRes)

	//dummyPackage.TriggerAWorkflow()

	updateTableWithNewMedia.UpdateThumbnailData()
}
