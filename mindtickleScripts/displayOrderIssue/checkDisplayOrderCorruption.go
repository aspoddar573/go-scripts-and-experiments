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

type CorruptedSeriesData struct {
	CompanyId string
	SeriesId  string
}

func FetchAllCorruptedSeriesForTrack() ([]CorruptedSeriesData, error) {
	var corruptedSeriesData []CorruptedSeriesData

	db := localSqlClient.GetSqlClient(LocalSchema)
	rows, err := db.Query("SELECT company_id, series_id FROM corrupted_series_display_order_data WHERE is_processed = 0 AND track = ?", TRACK)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var series_id string
		var company_id string
		if err := rows.Scan(&company_id, &series_id); err != nil {
			return corruptedSeriesData, err
		}
		corruptedSeriesData = append(corruptedSeriesData, CorruptedSeriesData{
			CompanyId: company_id,
			SeriesId:  series_id,
		})
	}
	if err = rows.Err(); err != nil {
		return corruptedSeriesData, err
	}

	defer db.Close()
	return corruptedSeriesData, nil
}

func EnhanceCorruptionDataForSeries(companyId string, seriesId string) error {
	corruptionInfo, err := CheckDisplayOrderCorruption(companyId, seriesId)
	if err != nil {
		fmt.Printf("Error occurred: %+v\n", err)
		return err
	}
	fmt.Println(corruptionInfo)
	db := localSqlClient.GetSqlClient(LocalSchema)

	_, err = db.Exec("UPDATE corrupted_series_display_order_data SET status = ?, message = ?, is_processed = 1 where company_id = ? and series_id = ?", corruptionInfo.StatusCode, corruptionInfo.ErrorMessage, companyId, seriesId)
	if err != nil {
		fmt.Printf("Failed to update info for series %s\n", seriesId)
		return err
	}

	defer db.Close()
	return nil
}

func EnhanceCorruptionDataForAllCompanies() error {
	db := localSqlClient.GetSqlClient(LocalSchema)

	corruptedSeriesData, err := FetchAllCorruptedSeriesForTrack()
	if err != nil {
		fmt.Printf("Error while fetching corrupted series data: %+v\n", err)
		return err
	}
	for idx, corruptedSeriesObj := range corruptedSeriesData {
		fmt.Printf("Adding data for company: %s and series: %s at pos %d of %d\n", corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId, idx+1, len(corruptedSeriesData))
		err := EnhanceCorruptionDataForSeries(corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId)
		if err != nil {
			fmt.Printf("error occurred at EnhanceCorruptionDataForSeries for company: %s and series: %s\n", corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId)
			return err
		}
	}

	defer db.Close()
	return nil
}
