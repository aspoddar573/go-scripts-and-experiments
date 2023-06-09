package localDashboardApi

type FetchAllProgramsForCompanyRequest struct {
	CompanyId string `json:"cname"`
}

type CheckDisplayOrderCorruptionForProgramRequest struct {
	CompanyId string `json:"cname"`
	SeriesId  string `json:"seriesId"`
}
