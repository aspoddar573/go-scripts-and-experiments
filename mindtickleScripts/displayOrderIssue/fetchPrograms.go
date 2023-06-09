package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/localDashboardApi"
	"goScriptsAndExperiments/mindtickleScripts/localSqlClient"
)

const LocalSchema = "display_order_issues"

func FetchAllPrograms(companyId string) (*localDashboardApi.FetchAllProgramsForCompanyResponse, error) {
	dashboardApiClient := localDashboardApi.GetDashboardApiClient()
	allPrograms, err := dashboardApiClient.FetchAllProgramsForCompany(&localDashboardApi.FetchAllProgramsForCompanyRequest{
		CompanyId: companyId,
	})
	if err == nil {
		fmt.Printf("Fetched %d number of programs for company %s\n", len(allPrograms.Data), companyId)
	}
	return allPrograms, err
}

func AddSeriesDataForCompany(companyId string) error {
	allPrograms, err := FetchAllPrograms(companyId)
	if err != nil {
		fmt.Printf("Error occurred: %+v\n", err)
		return err
	}
	db := localSqlClient.GetSqlClient(LocalSchema)

	fmt.Printf("Inserting %d entries for company %s\n", len(allPrograms.Data), companyId)
	for _, program := range allPrograms.Data {
		_, err = db.Exec("INSERT INTO series_display_order_status(company_id, series_id, track) VALUES(?, ?, ?)", companyId, program.Id, TRACK)
		if err != nil {
			fmt.Printf("Failed to inserted %s, %s and %s\n", companyId, program.Id, TRACK)
			return err
		}
	}
	fmt.Printf("Successfully inserted %d entries for company %s.\n", len(allPrograms.Data), companyId)

	_, err = db.Exec("UPDATE customer_company_info_table SET is_processed=1 where company_id = ? and track = ?", companyId, TRACK)
	if err != nil {
		fmt.Printf("Error while updating table for company: %s", companyId)
		return err
	}

	defer db.Close()
	return nil
}

func AddSeriesDataForCompanies(companyIds []string) error {

	for idx, companyId := range companyIds {
		fmt.Printf("Adding series data for company: %s at pos %d of %d\n", companyId, idx+1, len(companyIds))
		err := AddSeriesDataForCompany(companyId)
		if err != nil {
			return err
		}
	}

	return nil
}
