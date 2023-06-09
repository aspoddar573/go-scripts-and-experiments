package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/localSqlClient"
)

type CorruptedSeriesData struct {
	CompanyId string
	SeriesId  string
}

func FetchAllCorruptedSeriesForTrack() ([]CorruptedSeriesData, error) {
	var corruptedSeriesData []CorruptedSeriesData

	db := localSqlClient.GetSqlClient(LocalSchema)
	rows, err := db.Query("SELECT company_id, series_id FROM series_display_order_status WHERE status>200 AND track = ? order by company_id desc", TRACK)
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

	_, err = db.Exec("UPDATE series_display_order_status SET status = ?, message = ? where company_id = ? and series_id = ?", corruptionInfo.StatusCode, corruptionInfo.ErrorMessage, companyId, seriesId)
	if err != nil {
		fmt.Printf("Failed to update info for series %s\n", seriesId)
		return err
	}

	defer db.Close()
	return nil
}

func CorrectCorruptedDataForAllCompanies() error {
	db := localSqlClient.GetSqlClient(LocalSchema)

	corruptedSeriesData, err := FetchAllCorruptedSeriesForTrack()
	if err != nil {
		fmt.Printf("Error while fetching corrupted series data: %+v\n", err)
		return err
	}
	for idx, corruptedSeriesObj := range corruptedSeriesData {
		fmt.Printf("Correcting info for company: %s and series: %s at pos %d of %d\n", corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId, idx+1, len(corruptedSeriesData))
		correctionResponse, err := CorrectDisplayOrderCorruption(corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId)
		if err != nil || correctionResponse == nil || correctionResponse.StatusCode == 500 {
			fmt.Printf("Error occurred in correction for series: %s: %+v\n", corruptedSeriesObj.SeriesId, err)
		} else {
			fmt.Printf("Data corrected for series: %s with response %+v\n", corruptedSeriesObj.SeriesId, correctionResponse)
		}

		fmt.Printf("Adding data for company: %s and series: %s at pos %d of %d\n", corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId, idx+1, len(corruptedSeriesData))
		err = EnhanceCorruptionDataForSeries(corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId)
		if err != nil {
			fmt.Printf("error occurred at EnhanceCorruptionDataForSeries for company: %s and series: %s\n", corruptedSeriesObj.CompanyId, corruptedSeriesObj.SeriesId)
			return err
		}
	}

	defer db.Close()
	return nil
}
