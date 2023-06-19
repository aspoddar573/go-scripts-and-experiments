package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/fetchAllCustomerCompanies"
)

const TRACK = "prod-us"

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
	fmt.Println("####################################################################################################")
	fmt.Println("Adding series data for all companies\n\n")
	AddSeriesData()
	fmt.Println("####################################################################################################\n\n")

	fmt.Println("####################################################################################################")
	fmt.Println("Adding corruption data for all series\n\n")
	AddCorruptionData()
	fmt.Println("####################################################################################################\n\n")

	fmt.Println("####################################################################################################")
	fmt.Println("Correcting data for all corrupted series\n\n")
	CorrectDataForAllCompanies()
	fmt.Println("####################################################################################################\n\n")
}
