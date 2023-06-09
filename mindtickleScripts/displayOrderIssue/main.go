package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/fetchAllCustomerCompanies"
)

const TRACK = "integration"

func AddCompanyData() {

}

func AddSeriesData() {
	allUnprocessedCompanies, err := fetchAllCustomerCompanies.FetchAllUnprocessedCompanies(TRACK)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Fetched %d number of unprocessed companies!\n", len(allUnprocessedCompanies))

	err = AddSeriesDataForCompanies(allUnprocessedCompanies)
	if err != nil {
		return
	}
}

func AddCorruptionData() {
	allUnprocessedCompanies, err := fetchAllCustomerCompanies.FetchAllUnprocessedCompanies(TRACK)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Fetched %d number of unprocessed companies!\n", len(allUnprocessedCompanies))

	err = AddCorruptionDataForAllCompanies(allUnprocessedCompanies)
	if err != nil {
		return
	}
}

func EnhanceCorruptionDataCorruptedSeries() {
	err := EnhanceCorruptionDataForAllCompanies()
	if err != nil {
		return
	}
}

func main() {
	//AddSeriesData()
	//AddCorruptionData()
	//EnhanceCorruptionDataCorruptedSeries()
}
