package localDashboardApi

type FetchAllProgramsForCompanyResponse struct {
	Total int          `json:"total"`
	Data  []ProgramIds `json:"data"`
}

type ProgramIds struct {
	Id string `json:"id"`
}

type CheckDisplayOrderCorruptionForProgramResponse struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
}
