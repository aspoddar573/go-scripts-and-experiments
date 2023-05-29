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

	for _, program := range allPrograms.Data {
		_, err = db.Exec("INSERT INTO series_display_order_status(company_id, series_id, track) VALUES(?, ?, ?)", companyId, program.Id, TRACK)
		if err != nil {
			fmt.Printf("Failed to inserted %s, %s and %s\n", companyId, program.Id, TRACK)
			return err
		}
		fmt.Printf("Successfully inserted %s, %s and %s\n", companyId, program.Id, TRACK)
	}

	_, err = db.Exec("UPDATE customer_company_info_table SET is_processed=1 where company_id = ? and track = ?", companyId, TRACK)
	if err != nil {
		fmt.Printf("Error while updating table for company: %s", companyId)
		return err
	}

	return nil
}

func AddSeriesDataForCompanies(companyIds []string) error {

	for _, companyId := range companyIds {
		fmt.Printf("Adding series data for company: %s\n", companyId)
		err := AddSeriesDataForCompany(companyId)
		if err != nil {
			return err
		}
	}

	return nil
}
