// all constants string common to infra will be declared here
package infraconstant

type InfraConstant string

type MonitoringTag string

const (
	// This key is used for retrieving global context OperationId from request metadata.
	GlobalConstantId InfraConstant = "GLOBAL_CONTEXT_ID"

	OrgId InfraConstant = "ORG_ID"

	// This key is used for sending requestId to when an rpc is called
	CallerReqId InfraConstant = "callerReqId"

	// This key is used for sending requestId when an rpc is reponded back
	CalleeReqId InfraConstant = "calleeReqId"

	// This key used to retrieve requestId from request  context present in the service.
	ReqId InfraConstant = "reqId"
	// This key is used to retrieve response from headers metadata.
	ResponseId InfraConstant = "responseId"

	TagTenant         MonitoringTag = "tenant"
	TagResource       MonitoringTag = "resource"
	TagResourceName   MonitoringTag = "resource_name"
	TagProcessor      MonitoringTag = "processor"
	TagProcessorName  MonitoringTag = "processor_name"
	TagProcessorNameSpace MonitoringTag = "processor_name_space"
	TagNamespace 	  MonitoringTag	= "namespace"
)
