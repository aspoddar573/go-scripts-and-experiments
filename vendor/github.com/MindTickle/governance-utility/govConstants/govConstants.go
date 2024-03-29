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
	ENV_VAR_TRACK       string = "TRACK"
	DATADOG_APM_ENABLED string = "DATADOG_APM_ENABLED"
)

const (
	DEFAULT_AUTHORIZER_ID = "MT_ADMIN"
)

type WorkflowJobType string

const (
	CERTIFICATION_BULK_AWARD_CERTIFICATE       WorkflowJobType = "BulkAwardCertificate"
	CERTIFICATION_BULK_RESCIND_CERTIFICATE     WorkflowJobType = "BulkRescindCertificate"
	CERTIFICATION_EXPIRE_CERTIFICATE           WorkflowJobType = "ExpireCertificate"
	CERTIFICATION_DEACTIVATE_CERTIFICATIONS    WorkflowJobType = "DeactivateCertifications"
	CERTIFICATION_END_RECERTIFICATION_PERIOD   WorkflowJobType = "EndRecertificationPeriod"
	CERTIFICATION_START_RECERTIFICATION_PERIOD WorkflowJobType = "StartRecertificationPeriod"
	PUBLISH_CERTIFICATE_TEMPLATE               WorkflowJobType = "PublishCertificateTemplate"
	CREATE_MODULE_FROM_TEMPLATE                WorkflowJobType = "CreateModuleFromTemplate"
)
