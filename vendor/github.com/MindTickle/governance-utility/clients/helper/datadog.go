package helper

import (
	"context"
	"github.com/MindTickle/governance-utility/govConstants"
	"github.com/MindTickle/governance-utility/helper"
	"github.com/MindTickle/infracommon/constant/infraconstant"
	. "github.com/MindTickle/mt-go-logger/logger"
	"google.golang.org/grpc/metadata"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

func IsDatadogEnabled(ctx context.Context) bool {
	isDatadogEnabled := helper.GetEnv(govConstants.DATADOG_APM_ENABLED, "")
	if isDatadogEnabled == "true" {
		return true
	}
	Logger.Warnf(ctx, "gov-utility | IsDatadogEnabled | datadog is not enabled. ENV value : ", isDatadogEnabled)
	return false
}

func GetRequestIdFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	reqId := ctx.Value(string(infraconstant.ReqId))
	if reqId != nil {
		return reqId.(string)
	}
	reqId = ctx.Value(string(infraconstant.CallerReqId))
	if reqId != nil {
		return reqId.(string)
	}
	reqId = ctx.Value(string(infraconstant.GlobalConstantId))
	if reqId != nil {
		return reqId.(string)
	}
	reqId = ctx.Value(string(infraconstant.CalleeReqId))
	if reqId != nil {
		return reqId.(string)
	}
	reqId = ctx.Value("service-req-id")
	if reqId != nil {
		return reqId.(string)
	}
	return ""
}

func addReqIDTagsToSpan(ctx context.Context) {
	span, ok := tracer.SpanFromContext(ctx)
	if ok {
		reqId := GetRequestIdFromContext(ctx)
		if reqId != "" {
			span.SetTag(string(infraconstant.ReqId), reqId)
		}
	}
}

func InjectSpanIntoContextForRpc(ctx context.Context) context.Context {
	isDatadogEnabled := IsDatadogEnabled(ctx)
	if !isDatadogEnabled {
		return ctx
	}
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		Logger.Warnf(ctx, "gov-utility | failed to get span from context in InjectSpanIntoContextForRpc")
		return ctx
	}
	addReqIDTagsToSpan(ctx)
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		// we have to copy the metadata because its not safe to modify
		md = md.Copy()
	} else {
		md = metadata.MD{}
	}
	if err := tracer.Inject(span.Context(), MDCarrier(md)); err != nil {
		// in practice this error should never really happen
		Logger.Warnf(ctx, "gov-utility | failed to inject the span context into the gRPC metadata", err)
		return ctx
	}
	Logger.Debugf(ctx, "gov-utility | InjectSpanIntoContextForRpc process completed")
	return metadata.NewOutgoingContext(ctx, md)
}

func InjectSpanContextInReqHeaderForHTTPReq(ctx context.Context, reqHeader http.Header) {
	isDatadogEnabled := IsDatadogEnabled(ctx)
	if !isDatadogEnabled {
		return
	}
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		Logger.Warnf(ctx, "gov-utility | failed to get span from context in InjectSpanContextInReqHeaderForHTTPReq")
		return
	}
	// Inject the span Context in the Request headers
	err := tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(reqHeader))
	if err != nil {
		Logger.Warnf(ctx, "gov-utility | failed to inject the span context into the HTTP request header", err)
		return
	}
	Logger.Debugf(ctx, "gov-utility | InjectSpanContextInReqHeaderForHTTPReq process completed")
}
