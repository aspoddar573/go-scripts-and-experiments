package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/localDashboardApi"
	"goScriptsAndExperiments/mindtickleScripts/localSqlClient"
)

func CheckDisplayOrderCorruption(companyId string, seriesId string) (*localDashboardApi.CheckDisplayOrderCorruptionForProgramResponse, error) {
	dashboardApiClient := localDashboardApi.GetDashboardApiClient()
	corruptionInfo, err := dashboardApiClient.CheckDisplayOrderCorruptionForProgram(&localDashboardApi.CheckDisplayOrderCorruptionForProgramRequest{
		CompanyId: companyId,
		SeriesId:  seriesId,
	})
	return corruptionInfo, err
}

func CorrectDisplayOrderCorruption(companyId string, seriesId string) (*localDashboardApi.CheckDisplayOrderCorruptionForProgramResponse, error) {
	dashboardApiClient := localDashboardApi.GetDashboardApiClient()
	correctionResponse, err := dashboardApiClient.CorrectDisplayOrderCorruptionForProgram(&localDashboardApi.CheckDisplayOrderCorruptionForProgramRequest{
		CompanyId: companyId,
		SeriesId:  seriesId,
	})
	return correctionResponse, err
}

func AddCorruptionDataForSeries(companyId string, seriesId string) error {
	corruptionInfo, err := CheckDisplayOrderCorruption(companyId, seriesId)
	if err != nil {
		fmt.Printf("Error occurred: %+v\n", err)
		return err
	}
	db := localSqlClient.GetSqlClient(LocalSchema)

	_, err = db.Exec("UPDATE series_display_order_status SET status = ?, message = ?, is_processed = 1 where company_id = ? and series_id = ?", corruptionInfo.StatusCode, corruptionInfo.ErrorMessage, companyId, seriesId)
	if err != nil {
		fmt.Printf("Failed to update info for series %s\n", seriesId)
		return err
	}

	defer db.Close()
	return nil
}

func FetchAllUnprocessedSeriesForCompany(companyId string) ([]string, error) {
	var allUnprocessedSeries []string

	db := localSqlClient.GetSqlClient(LocalSchema)
	rows, err := db.Query("SELECT series_id FROM series_display_order_status WHERE company_id = ? and is_processed = 0", companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var series_id string
		if err := rows.Scan(&series_id); err != nil {
			return allUnprocessedSeries, err
		}
		allUnprocessedSeries = append(allUnprocessedSeries, series_id)
	}
	if err = rows.Err(); err != nil {
		return allUnprocessedSeries, err
	}

	defer db.Close()
	return allUnprocessedSeries, nil
}

func AddCorruptionDataForCompany(companyId string) error {
	allUnprocessedSeries, err := FetchAllUnprocessedSeriesForCompany(companyId)
	if err != nil {
		fmt.Printf("Error fetching unprocessed series for company: %s.\n", companyId)
		return err
	}

	for idx, seriesId := range allUnprocessedSeries {
		fmt.Printf("Adding data for series: %s at pos %d of %d\n", seriesId, idx+1, len(allUnprocessedSeries))
		err := AddCorruptionDataForSeries(companyId, seriesId)
		if err != nil {
			fmt.Printf("Error occurred while adding data for series: %s for company: %s.\n", seriesId, companyId)
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func AddCorruptionDataForAllCompanies(allUnprocessedCompanies []string) error {
	db := localSqlClient.GetSqlClient(LocalSchema)

	for idx, companyId := range allUnprocessedCompanies {
		fmt.Printf("Adding data for company: %s at pos %d of %d\n", companyId, idx+1, len(allUnprocessedCompanies))
		err := AddCorruptionDataForCompany(companyId)
		if err != nil {
			fmt.Printf("error occurred at AddCorruptionDataForAllCompanies for company: %s\n", companyId)
			return err
		}

		_, err = db.Exec("UPDATE customer_company_info_table SET is_processed=1 where company_id = ? and track = ?", companyId, TRACK)
		if err != nil {
			fmt.Printf("Error while updating table for company: %s", companyId)
			return err
		}
	}

	defer db.Close()
	return nil
}
