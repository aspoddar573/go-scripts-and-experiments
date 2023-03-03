package httpRequests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//type CompanyObject

func GetSeries() {
	var companyID int64 = 1262345569886290823
	//bodyBytes, err := json.Marshal(map[string]string{
	//	"company": "1262345569886290823",
	//})
	//body := bytes.NewBuffer(bodyBytes)
	//resp, err := http.Post("http://ce.internal.staging.mindtickle.com/company/settings/get", "application/json", body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	postReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://ce-governance.internal.staging.mindtickle.com/series/1503313078281523534?company=1453686784317271186", nil)
	postReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(postReq)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var sbs map[string]interface{} = make(map[string]interface{})
	json.Unmarshal(res, &sbs)
	v2, _ := json.Marshal(sbs["object"])
	json.Unmarshal(v2, &sbs)
	companyIdString := strconv.FormatInt(companyID, 10)
	v2, _ = json.Marshal(sbs[companyIdString])
	json.Unmarshal(v2, &sbs)
	fmt.Println(sbs["displayName"])
	fmt.Println(resp)
	fmt.Println(err)
}

func GetCompanySettings() {
	var companyID int64 = 1262345569886290823
	bodyBytes, err := json.Marshal([]int64{companyID})
	body := bytes.NewBuffer(bodyBytes)
	//resp, err := http.Post("http://ce.internal.staging.mindtickle.com/company/settings/get", "application/json", body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	postReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://ce.internal.staging.mindtickle.com/company/settings/get", body)
	postReq.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(postReq)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var sbs map[string]interface{} = make(map[string]interface{})
	json.Unmarshal(res, &sbs)
	v2, _ := json.Marshal(sbs["object"])
	json.Unmarshal(v2, &sbs)
	companyIdString := strconv.FormatInt(companyID, 10)
	v2, _ = json.Marshal(sbs[companyIdString])
	json.Unmarshal(v2, &sbs)
	fmt.Println(sbs["displayName"])
	fmt.Println(resp)
	fmt.Println(err)
}
