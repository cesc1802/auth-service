package logger

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/rs/xid"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *Logger = &Logger{}

type Logger struct {
	*zap.Logger
	options *options
}

func Init(opts ...applyOption) *Logger {
	options := defaultOption()
	for _, apply := range opts {
		apply(options)
	}

	return initLogger(options)

}

func New(name string, opts ...zap.Option) *Logger {
	return &Logger{
		Logger: globalLogger.Named(name).WithOptions(opts...),
	}
}

func (logger *Logger) New(opts ...zap.Option) *Logger {
	return &Logger{
		Logger: logger.WithOptions(opts...),
	}
}
func (logger *Logger) WithCaller() *Logger {
	var clone = logger.clone()
	clone.Logger = clone.WithOptions(zap.AddCaller())
	return clone
}

func (logger *Logger) WithSkipCaller(skip int) *Logger {
	var clone = logger.clone()
	clone.Logger = clone.WithOptions(zap.AddCallerSkip(skip))
	return clone
}

func (logger *Logger) WithRequestID(id ...string) *Logger {
	var clone = logger.clone()
	var reqID = xid.New().String()
	if len(id) > 0 {
		reqID = id[0]
	}
	clone.Logger = clone.With(zap.String(RequestIDKey, reqID))
	return clone
}

func (logger *Logger) WithUserID(id string) *Logger {
	var clone = logger.clone()
	clone.Logger = clone.With(zap.String(UserIDKey, id))
	return clone
}

func (logger *Logger) WithRequestResponse(req interface{}, res interface{}) *Logger {
	var clone = logger.clone()
	clone.Logger = clone.With(zap.Any(ApiReqDataKey, req), zap.Any(ApiResDataKey, res))
	return clone
}

func (logger *Logger) WithFields(fields ...zap.Field) *Logger {
	var clone = logger.clone()
	clone.Logger = clone.With(fields...)
	return clone
}

func (logger *Logger) WithContext(ctx context.Context) *Logger {
	var clone = logger.clone()
	if id := ctx.Value(RequestIDKey); id != nil {
		if reqID, ok := id.(string); ok {
			return clone.WithRequestID(reqID)
		}
	}
	return clone
}

func (logger *Logger) WithRequestInfo(info *RequestInfo) *Logger {
	var clone = logger.clone()
	if info == nil {
		return clone
	}

	var fields []zap.Field
	if info.Method != "" {
		fields = append(fields, zap.String(ApiMethodKey, info.Method))
	}

	if info.ReqID != "" {
		fields = append(fields, zap.String(RequestIDKey, info.ReqID))
	}

	if info.RefErrorID != "" {
		fields = append(fields, zap.String(ErrorRefIDKey, info.Method))
	}

	if info.URI != "" {
		fields = append(fields, zap.String(ApiURIKey, info.URI))
	}

	if info.StatusCode > 0 {
		fields = append(fields, zap.Int(ApiStatusCodeKey, info.StatusCode))
	}

	if info.Size > 0 {
		fields = append(fields, zap.Int64(ApiRequestSizeKey, info.Size))
	}

	if info.Latency > 0 {
		fields = append(fields, zap.Duration(ApiLatencyKey, info.Latency))
	}

	if info.UserAgent != "" {
		fields = append(fields, zap.String(ApiUserAgentKey, info.UserAgent))
	}

	if info.RemoteIP != "" {
		fields = append(fields, zap.String(RemoteIPKey, info.RemoteIP))
	}

	if info.Host != "" {
		fields = append(fields, zap.String(ApiHostKey, info.Host))
	}

	if info.URL != "" {
		var url = fmt.Sprintf("%s %s", info.Method, info.URL)
		fields = append(fields, zap.String(ApiURLKey, url))
	}

	clone.Logger = clone.With(fields...)
	return clone
}

func (logger *Logger) WithHTTPRequest(req *http.Request) *Logger {
	var clone = logger.clone()
	if req == nil {
		return clone
	}

	var fields = []zap.Field{}

	if req.Method != "" {
		fields = append(fields, zap.String(ApiMethodKey, req.Method))
	}

	if req.Response != nil && req.Response.StatusCode > 0 {
		fields = append(fields, zap.Int(ApiStatusCodeKey, req.Response.StatusCode))
	}

	if req.URL.Host != "" {
		fields = append(fields, zap.String(ApiHostKey, req.URL.Host))
	}

	if req.RequestURI != "" {
		fields = append(fields, zap.String(ApiURIKey, req.RequestURI))
	} else {
		fields = append(fields, zap.String(ApiURIKey, req.URL.Path))
	}

	if len(req.Header) > 0 {
		fields = append(fields, zap.Any(ApiHeaderKey, "helper.FormatHeaders(req.Header)"))
	}

	if req.URL.String() != "" {
		var url = fmt.Sprintf("%s %s", req.Method, req.URL.String())
		fields = append(fields, zap.String(ApiURLKey, url))
	}

	clone.Logger = clone.With(fields...)
	return clone
}

func (logger *Logger) WithRestyRequestError(r *resty.Request, err error) *Logger {
	var clone = logger.clone()
	var fields = []zap.Field{}

	if r.Body != nil {
		fields = append(fields, zap.Any(ApiReqDataKey, r.Body))
	}

	if r.FormData != nil {
		fields = append(fields, zap.Any(ApiReqFormDataKey, r.FormData))
	}

	if r.QueryParam != nil {
		fields = append(fields, zap.Any(ApiReqParamsKey, r.QueryParam))
	}

	if r.Result != nil {
		fields = append(fields, zap.Any(ApiResDataKey, r.Result))
	} else {
		if r.Error != nil {
			fields = append(fields, zap.Any(ApiResDataKey, r.Error))
		}
	}

	if r.Method != "" {
		fields = append(fields, zap.String(ApiMethodKey, r.Method))
	}

	if r.RawRequest != nil {
		if r.RawRequest.Host != "" {
			fields = append(fields, zap.String(ApiHostKey, r.RawRequest.Host))
		}

		if r.RawRequest.RequestURI != "" {
			fields = append(fields, zap.String(ApiURIKey, r.RawRequest.RequestURI))
		} else {
			fields = append(fields, zap.String(ApiURIKey, r.RawRequest.URL.Path))
		}

		if len(r.RawRequest.Header) > 0 {
			fields = append(fields, zap.Any(ApiHeaderKey, "helper.FormatHeaders(r.RawRequest.Header)"))
		}

		if r.URL != "" {
			var url = fmt.Sprintf("%s %s", r.Method, r.URL)
			fields = append(fields, zap.String(ApiURLKey, url))
		}

	}

	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	clone.Logger = clone.With(fields...)
	return clone
}

func (logger *Logger) WithRestyResponse(r *resty.Response) *Logger {
	var clone = logger.clone()
	var fields = []zap.Field{}

	if r.Request.Body != nil {
		fields = append(fields, zap.Any(ApiReqDataKey, r.Request.Body))
	}

	if r.Request.FormData != nil {
		fields = append(fields, zap.Any(ApiReqFormDataKey, r.Request.FormData))
	}

	if r.Request.QueryParam != nil {
		fields = append(fields, zap.Any(ApiReqParamsKey, r.Request.QueryParam))
	}

	if r.Request.Result != nil {
		fields = append(fields, zap.Any(ApiResDataKey, r.Request.Result))
	} else {
		if r.Request.Error != nil {
			fields = append(fields, zap.Any(ApiResDataKey, r.Request.Error))
		}
	}

	if r.Request.Method != "" {
		fields = append(fields, zap.String(ApiMethodKey, r.Request.Method))
	}

	if r.StatusCode() > 0 {
		fields = append(fields, zap.Int(ApiStatusCodeKey, r.StatusCode()))
	}

	if r.Request.RawRequest != nil {
		if r.Request.RawRequest.Host != "" {
			fields = append(fields, zap.String(ApiHostKey, r.Request.RawRequest.Host))
		}

		if r.Request.RawRequest.RequestURI != "" {
			fields = append(fields, zap.String(ApiURIKey, r.Request.RawRequest.RequestURI))
		} else {
			fields = append(fields, zap.String(ApiURIKey, r.Request.RawRequest.URL.Path))
		}

		if len(r.Request.RawRequest.Header) > 0 {
			fields = append(fields, zap.Any(ApiHeaderKey, "helper.FormatHeaders(r.Request.RawRequest.Header)"))
		}

		if r.Request.URL != "" {
			var url = fmt.Sprintf("%s %s", r.Request.Method, r.Request.URL)
			fields = append(fields, zap.String(ApiURLKey, url))
		}
	}

	clone.Logger = clone.With(fields...)
	return clone
}

func (logger *Logger) InfoAny(data interface{}, msg ...string) {
	var message = ""
	if len(msg) > 0 {
		message = msg[0]
	}

	logger.WithOptions(zap.AddCallerSkip(1)).Info(message, zap.Any(DataKey, data))
}

func (logger *Logger) Infof(format string, data ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Sugar().Infof(format, data...)
}

func (logger *Logger) DebugAny(data interface{}, msg ...string) {
	var message = ""
	if len(msg) > 0 {
		message = msg[0]
	}

	logger.WithOptions(zap.AddCallerSkip(1)).Debug(message, zap.Any(DataKey, data))
}

func (logger *Logger) Debugf(format string, data ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Sugar().Debugf(format, data...)
}

func (logger *Logger) ErrorAny(data interface{}, msg ...string) {
	var message = ""
	if len(msg) > 0 {
		message = msg[0]
	}

	switch value := data.(type) {
	case error:
		//msg, stack := errs.ParseErisJSON(value)
		logger.WithOptions(zap.AddCallerSkip(1)).Error("msg", zap.Any(StackKey, value))
	default:
		logger.WithOptions(zap.AddCallerSkip(1)).Error(message, zap.Any(DataKey, data))
	}

}

func (logger *Logger) Errorf(format string, data ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Sugar().Errorf(format, data...)
}

func (logger *Logger) WarnAny(data interface{}, msg ...string) {
	var message = ""
	if len(msg) > 0 {
		message = msg[0]
	}

	logger.WithOptions(zap.AddCallerSkip(1)).Warn(message, zap.Any(DataKey, data))
}

func (logger *Logger) Warnf(format string, data ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Sugar().Warnf(format, data...)
}

func (logger *Logger) PanicAny(data interface{}, msg ...string) {
	var message = ""
	if len(msg) > 0 {
		message = msg[0]
	}

	logger.WithOptions(zap.AddCallerSkip(1)).Panic(message, zap.Any(DataKey, data))
}

func (logger *Logger) clone() *Logger {
	copy := *logger
	return &copy
}

func initLogger(options *options) *Logger {
	var encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "module",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	outputPaths, errorOutputPaths, err := options.configDailyRotate()
	if err != nil {
		panic(fmt.Sprintf("Config log errors: %+v", err))
	}

	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(options.logLevel),
		Development:      options.debug,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorOutputPaths,
	}

	if options.debug {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	if options.logConsole {
		loggerConfig.OutputPaths = append(loggerConfig.OutputPaths, "stdout")
		loggerConfig.ErrorOutputPaths = append(loggerConfig.ErrorOutputPaths, "stdout")
	}

	if !options.logJSON {
		loggerConfig.Encoding = "console"
	}

	zapLogger, err := loggerConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("build zap logger from config error: %v", err))
	}

	globalLogger = &Logger{
		Logger:  zapLogger,
		options: options,
	}

	return globalLogger
}
