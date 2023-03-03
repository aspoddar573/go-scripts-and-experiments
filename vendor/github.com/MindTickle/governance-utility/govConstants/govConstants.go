package govConstants

const (
	TEAM_NAME_SHORT     string = "gov"
	CORRELATION_ID      string = "correlation_id" // common id for a request to be passed on from each request
	CALLER_REQ_ID       string = "caller_req_id"  // request id of the call from upstream service
	CALLEE_REQ_ID       string = "callee_req_id"  // request id generated by the service for the call
	COMPANY_ID          string = "company_id"
	ORG_ID              string = "org_id"
	AUTHORIZER_ID       string = "authorizer_id"
	HOSTNAME            string = "HOSTNAME"
	APPNAME             string = "APPLICATION"
	PROJECT             string = "governance"
	TICKLE_DB_NAMESPACE string = "governance"
	REQ_ID              string = "reqId"
	SESSION_KEY         string = "session_key"
	CONTEXT_TYPE        string = "context_type"
)

const (
	DEFAULT_AUTHORIZER_ID = "MT_ADMIN"
)
