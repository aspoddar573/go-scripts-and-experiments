package main

import (
	"fmt"
	"github.com/MindTickle/governance-utility/helper"
	"goScriptsAndExperiments/mindtickleScripts/templates/release/fetchAllActiveCompanies"
)

const (
	TRACK   = "integration"
	DIVIDER = "================================================================================="
)

func main() {
	allCompanyObjects, err := fetchAllActiveCompanies.FetchAllCompaniesForTrack(TRACK)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Fetched %d company-ids!\n", len(allCompanyObjects))

	//ceClient := setTemplatesV2Flag.GetContentEngineClient(TRACK)
	//setFlag := true

	for idx, companyObject := range allCompanyObjects {
		fmt.Println(DIVIDER)
		fmt.Printf("Calling content-engine to set 'templatesV2Enabled' for companyId: %d at idx: %d\n", companyObject.CompanyId, idx)

		//err = setTemplatesV2Flag.SetTemplatesV2FlagForCompany(ceClient, helper.Int64ToDecimalStr(companyObject.CompanyId), setFlag)
		fmt.Println(companyObject.CompanyId)
		fmt.Printf("%s\n", helper.Int64ToDecimalStr(companyObject.CompanyId))

		if err != nil {
			fmt.Printf("Failed to set 'templatesV2Enabled' for companyId: %d at idx: %d\n", companyObject.CompanyId, idx)
			fmt.Println(err)
		} else {
			fmt.Printf("Successfully set 'templatesV2Enabled' for companyId: %d at idx: %d\n", companyObject.CompanyId, idx)
		}
		fmt.Println(DIVIDER)

	}
}
