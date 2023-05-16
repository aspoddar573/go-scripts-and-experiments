package constants

var TracingTags = []string{ParentSpanIdTag, SampledTag, SpanIdTag, TraceIdTag, RequestIdTag, TenantIdTag, ServiceNameTag, CompanyIdTag, CompanyTypeTag}

const (
	ParentSpanIdTag string = "x-b3-parentspanid"
	SampledTag      string = "x-b3-sampled"
	SpanIdTag       string = "x-b3-spanid"
	TraceIdTag      string = "x-b3-traceid"
	RequestIdTag    string = "x-request-id"
	TenantIdTag     string = "x-tenant-id"
	ServiceNameTag  string = "x-service-name"
	CompanyIdTag    string = "x-company-id"
	CompanyTypeTag  string = "x-company-type"
)

const ServiceReqIdTag = "service-req-id"
const DefaultLogMsgSize int = 10000

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Fatal = "fatal"
	Audit = "audit"
)
