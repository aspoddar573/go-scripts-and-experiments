package localDashboardApi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	. "github.com/MindTickle/mt-go-logger/logger"
	"io/ioutil"
	"net/http"
	"time"
)

const LocalDashboardApiUrl = "http://localhost:8010"

type DashboardApiClientImpl struct {
	url     string
	timeout time.Duration
}

var dashboardApiRestClient DashboardApiClientImpl

func GetDashboardApiClient() DashboardApiClientImpl {
	dashboardApiRestClient = DashboardApiClientImpl{
		url:     LocalDashboardApiUrl,
		timeout: 10 * time.Second,
	}
	return dashboardApiRestClient
}

func (client DashboardApiClientImpl) FetchAllProgramsForCompany(req *FetchAllProgramsForCompanyRequest) (*FetchAllProgramsForCompanyResponse, error) {
	ctx := context.Background()

	jsonBytes, err := json.Marshal(req)
	postReq, err := http.NewRequestWithContext(ctx, http.MethodPost, client.url+"/api/v2/webhook/dashboard/programs", bytes.NewBuffer(jsonBytes))
	postReq.Header.Set("Content-Type", "application/json")
	postReq.Header.Set("x-token", "localkey")
	postReq.Header.Set("cname", req.CompanyId)
	postReq.Header.Set("user", "abcd")
	postReq.Header.Set("workflow-id", "abcd")
	postReq.Header.Set("domain-base", "abcd")

	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		fmt.Printf("Fetched all Programs for company: %d\n", req.CompanyId)
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error while reading response body: %+v\n", err)
		return nil, err
	}
	var responseObject = FetchAllProgramsForCompanyResponse{}

	err = json.Unmarshal(res, &responseObject)
	if err != nil {
		Logger.WithTag("error", err).Errorf(ctx, "Error while unmarshalling response object")
		return nil, err
	}
	return &responseObject, nil
}

func (client DashboardApiClientImpl) CheckDisplayOrderCorruptionForProgram(req *CheckDisplayOrderCorruptionForProgramRequest) (*CheckDisplayOrderCorruptionForProgramResponse, error) {
	ctx := context.Background()

	jsonBytes, err := json.Marshal(req)
	postReq, err := http.NewRequestWithContext(ctx, http.MethodPost, client.url+"/api/v2/webhook/identifyCorruptDisplayOrders", bytes.NewBuffer(jsonBytes))
	postReq.Header.Set("Content-Type", "application/json")
	postReq.Header.Set("x-token", "localkey")
	postReq.Header.Set("cname", req.CompanyId)
	postReq.Header.Set("user", "abcd")
	postReq.Header.Set("workflow-id", "abcd")
	postReq.Header.Set("domain-base", "abcd")

	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		fmt.Printf("Fetched all Programs for company: %d\n", req.CompanyId)
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error while reading response body: %+v\n", err)
		return nil, err
	}
	var responseObject = CheckDisplayOrderCorruptionForProgramResponse{}

	err = json.Unmarshal(res, &responseObject)
	if err != nil {
		Logger.WithTag("error", err).Errorf(ctx, "Error while unmarshalling response object")
		return nil, err
	}
	return &responseObject, nil
}
