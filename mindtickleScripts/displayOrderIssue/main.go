package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/fetchAllCustomerCompanies"
)

const TRACK = "integration"

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
	allUnprocessedCompanies, err := fetchAllCustomerCompanies.FetchAllUnprocessedSeriesCompanies(TRACK)
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

func CorrectDataForAllCompanies() {
	err := CorrectCorruptedDataForAllCompanies()
	if err != nil {
		return
	}
}

func main() {
	AddSeriesData()
	AddCorruptionData()
	CorrectDataForAllCompanies()
}
