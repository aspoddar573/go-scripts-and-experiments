package fetchAllCustomerCompanies

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MindTickle/governance-utility/helper"
	"github.com/MindTickle/governance-utility/storage/tickledb"
	"github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"goScriptsAndExperiments/mindtickleScripts/localSqlClient"
	"goScriptsAndExperiments/mindtickleScripts/templates/programTemplates-v1.0/sqlClient"
	"goScriptsAndExperiments/mindtickleScripts/templates/utils"
)

const (
	NAMESPACE   = "platform"
	LocalSchema = "display_order_issues"
)

type CompanyObject struct {
	CompanyId string `json:"company_id"`
}

func FetchAllCompaniesForTrack(track string) ([]CompanyObject, error) {
	var allCompaniesForTrack []CompanyObject
	allCompanyIds := map[string]bool{}

	ctx := context.Background()
	targetSqlStoreClient := sqlClient.GetTickleDBClient(track)

	searchRes, err := targetSqlStoreClient.Search(ctx, &tickleDbSqlStore.SearchRequest{
		RequestContext: utils.GetRequestContext(0, track, NAMESPACE),
		SqlStatement:   "SELECT ci.company_id as company_id FROM company_info ci where (ci.company_type = 'CUSTOMER')",
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, row := range searchRes.Rows {
		rowBytes, err := tickledb.ConvertRowToDbModel(ctx, row)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		rowData := &CompanyObject{}
		err = json.Unmarshal(rowBytes, rowData)
		if err != nil {
			fmt.Println(rowBytes)
			fmt.Println(err)
			return nil, err
		}
		allCompanyIds[rowData.CompanyId] = true
	}

	for stringCname, intCname := range helper.CompanyIdsMap {
		strIntCname := helper.Int64ToDecimalStr(intCname)
		if _, ok := allCompanyIds[strIntCname]; ok {
			delete(allCompanyIds, strIntCname)
			allCompanyIds[stringCname] = true
		}
	}

	for cname, _ := range allCompanyIds {
		allCompaniesForTrack = append(allCompaniesForTrack, CompanyObject{
			CompanyId: cname,
		})
	}

	return allCompaniesForTrack, nil
}

func UpdateCompanyInfoInTable(TRACK string) error {
	allCompanyData, err := FetchAllCompaniesForTrack(TRACK)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Fetched %d number of companies!\n", len(allCompanyData))

	db := localSqlClient.GetSqlClient(LocalSchema)

	for _, val := range allCompanyData {
		_, err := db.Exec("INSERT INTO customer_company_info_table(company_id, track) VALUES(?, ?)", val.CompanyId, TRACK)
		if err != nil {
			fmt.Printf("Failed to inserted %s and %s\n", val.CompanyId, TRACK)
		}
		fmt.Printf("Successfully inserted %s and %s\n", val.CompanyId, TRACK)
	}

	return nil
}

func FetchAllUnprocessedCompanies(TRACK string) ([]string, error) {
	var allUnprocessedCompanies []string

	db := localSqlClient.GetSqlClient(LocalSchema)
	rows, err := db.Query("SELECT company_id FROM customer_company_info_table WHERE track = ? and is_processed = 0", TRACK)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var company_id string
		if err := rows.Scan(&company_id); err != nil {
			return allUnprocessedCompanies, err
		}
		allUnprocessedCompanies = append(allUnprocessedCompanies, company_id)
	}
	if err = rows.Err(); err != nil {
		return allUnprocessedCompanies, err
	}
	return allUnprocessedCompanies, nil
}
