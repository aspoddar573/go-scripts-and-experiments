package main

import (
	"fmt"
	"github.com/MindTickle/governance-utility/helper"
	"goScriptsAndExperiments/mindtickleScripts/templates/release/fetchAllActiveCompanies"
	"goScriptsAndExperiments/mindtickleScripts/templates/release/setTemplatesV2Flag"
)

const (
	TRACK   = "integration"
	DIVIDER = "================================================================================="
)

func main1() {
	fmt.Printf("Running the script for the given track: %s\n", TRACK)

	allCompanyObjects, err := fetchAllActiveCompanies.FetchAllCompaniesForTrack(TRACK)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Fetched %d company-ids!\n", len(allCompanyObjects))

	ceClient := setTemplatesV2Flag.GetContentEngineClient(TRACK)
	setFlag := true

	for idx, companyObject := range allCompanyObjects {
		fmt.Println(DIVIDER)
		fmt.Printf("Calling content-engine to set 'templatesV2Enabled' for companyId: %d at idx: %d\n", companyObject.CompanyId, idx)

		err := setTemplatesV2Flag.SetTemplatesV2FlagForCompany(ceClient, helper.Int64ToDecimalStr(companyObject.CompanyId), setFlag)

		if err != nil {
			fmt.Printf("Failed to set 'templatesV2Enabled' for companyId: %d at idx: %d\n", companyObject.CompanyId, idx)
			fmt.Println(err)
		} else {
			fmt.Printf("Successfully set 'templatesV2Enabled' for companyId: %d at idx: %d\n", companyObject.CompanyId, idx)
		}
		fmt.Println(DIVIDER)

	}
}

func main2() {
	fmt.Printf("Running the script for the given track: %s\n", TRACK)

	ceClient := setTemplatesV2Flag.GetContentEngineClient(TRACK)
	setFlag := true

	idx := 0
	for key := range helper.CompanyIdsMap {
		fmt.Println(DIVIDER)
		fmt.Printf("Calling content-engine to set 'templatesV2Enabled' for companyId: %d at idx: %d\n", key, idx)

		err := setTemplatesV2Flag.SetTemplatesV2FlagForCompany(ceClient, key, setFlag)

		if err != nil {
			fmt.Printf("Failed to set 'templatesV2Enabled' for companyId: %d at idx: %d\n", key, idx)
			fmt.Println(err)
		} else {
			fmt.Printf("Successfully set 'templatesV2Enabled' for companyId: %d at idx: %d\n", key, idx)
		}
		fmt.Println(DIVIDER)

		idx++
	}
}

func main() {
	main1()
}
