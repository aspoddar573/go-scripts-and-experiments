package main

import (
	"fmt"
	"goScriptsAndExperiments/mindtickleScripts/localDashboardApi"
	"time"
)

func AmolIssue() {
	currTime := time.Now()
	fmt.Printf("\n\n\nStarting run: %=v\n\n\n", currTime)

	companyId := "1407275718740760934"
	dashboardApiClient := localDashboardApi.GetDashboardApiClient()

	fmt.Printf("\n\n\nFetching Templates:\n\n\n")
	allPrograms, err := dashboardApiClient.FetchAllProgramsForCompany(&localDashboardApi.FetchAllProgramsForCompanyRequest{
		CompanyId: companyId,
	})
	if err != nil {
		fmt.Printf("error occcurred while fetching programs: %+v\n", err)
	} else {
		fmt.Printf("fetched all programs for company: %+v\n", allPrograms)
	}

	fmt.Printf("\n\n\nEvaluating corruption info:\n\n\n")
	for _, programId := range allPrograms.Data {
		corruptionInfo, err := dashboardApiClient.CheckDisplayOrderCorruptionForProgram(&localDashboardApi.CheckDisplayOrderCorruptionForProgramRequest{
			CompanyId: companyId,
			SeriesId:  programId.Id,
		})
		if err != nil {
			fmt.Printf("error occcurred(%+v) while fetching corruption info for series: %s\n", programId.Id, err)
		} else {
			fmt.Printf("Corruption info %+v found for series id: %s\n", corruptionInfo, programId.Id)
		}
	}
}

func main() {
	AmolIssue()
}
