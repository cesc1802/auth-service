package logger

import (
	"time"

	"go.uber.org/zap"
)

const (
	RequestIDKey       = "request_id"
	UserIDKey          = "user_id"
	ApiLatencyKey      = "api_latency"
	ApiMethodKey       = "api_method"
	ApiResponseKey     = "api_response_size"
	ApiUserAgentKey    = "api_user_agent"
	ApiURIKey          = "api_uri"
	ApiURLKey          = "api_url"
	ApiHeaderKey       = "api_header"
	ApiHostKey         = "api_host"
	ApiRequestSizeKey  = "api_request_size"
	ApiReqDataKey      = "api_request_data"
	ApiReqFormDataKey  = "api_request_form_data"
	ApiReqParamsKey    = "api_request_params"
	ApiResDataKey      = "api_response_data"
	RemoteIPKey        = "remote_ip"
	ErrorRefIDKey      = "error_ref_id"
	ApiStatusCodeKey   = "api_status_code"
	DataKey            = "data"
	StackKey           = "stack"
	TaskStateKey       = "task_state"
	TaskAttemptsKey    = "task_attempts"
	TaskLatencyKey     = "task_latency"
	HeaderRequestIDKey = "X-Request-ID"
	HeaderXDVSignature = "X-DV-Signature"
)

type RequestInfo struct {
	ReqID      string
	UserID     string
	Method     string
	Host       string
	URI        string
	RefErrorID string
	StatusCode int
	RemoteIP   string
	UserAgent  string
	Latency    time.Duration
	Size       int64
	URL        string

	Header map[string]string
}

func (reqInfo *RequestInfo) ToZapFields(startTime ...time.Time) (fields []zap.Field) {
	fields = []zap.Field{
		zap.String(RequestIDKey, reqInfo.ReqID),
		zap.String(RemoteIPKey, reqInfo.RemoteIP),
		zap.String(ApiMethodKey, reqInfo.Method),
		zap.String(ApiHostKey, reqInfo.Host),
		zap.String(ApiURIKey, reqInfo.URI),
		zap.Int(ApiStatusCodeKey, reqInfo.StatusCode),
		zap.Int64(ApiRequestSizeKey, int64(reqInfo.Size)),
		zap.String(ApiUserAgentKey, reqInfo.UserAgent),
		zap.String(ApiURLKey, reqInfo.URL),
		zap.Any(ApiHeaderKey, reqInfo.Header),
		zap.Any(UserIDKey, reqInfo.UserID),
	}
	if len(startTime) > 0 {
		fields = append(fields, zap.Duration(ApiLatencyKey, time.Since(startTime[0])))
	}
	return fields
}
