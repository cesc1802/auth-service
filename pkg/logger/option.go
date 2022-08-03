package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
	"os"
	"time"
)

type applyOption func(o *options)

type options struct {
	logDir            string
	logFilename       string
	logConsole        bool
	logJSON           bool
	logLevel          zapcore.Level
	logFileMaxSize    int
	logFileMaxBackups int
	logFileMaxAge     int
	logFileCompress   bool
	timeLocation      *time.Location
	debug             bool
}

func defaultOption() *options {
	loc, _ := time.LoadLocation("Asia/Singapore")
	return &options{
		logDir:            "logs",
		logLevel:          zapcore.InfoLevel,
		debug:             true,
		logJSON:           true,
		logFileMaxSize:    1024 * 2, // 4GB
		logFileMaxBackups: 2,
		logFileMaxAge:     2, // days
		logFileCompress:   false,
		logFilename:       time.Now().In(loc).Format("2006-01-02.json"),
		timeLocation:      loc,
	}
}

func WithLogLevel(level zapcore.Level) applyOption {
	return func(o *options) {
		o.logLevel = level
	}
}

func WithLogDir(logDir string) applyOption {
	return func(o *options) {
		if logDir == "" {

			return
		}

		o.logDir = logDir
	}
}

func WithLogFilename(filename string) applyOption {
	return func(o *options) {
		if filename == "" {

			return
		}

		o.logFilename = filename
	}
}

func WithConsole(enable bool) applyOption {
	return func(o *options) {
		o.logConsole = enable
	}
}

func WithJSON(enable bool) applyOption {
	return func(o *options) {
		o.logJSON = enable
	}
}

func WithDebug(isDebug bool) applyOption {
	return func(o *options) {
		o.debug = isDebug
	}
}

func WithLogFileMaxSize(value int) applyOption {
	return func(o *options) {
		o.logFileMaxSize = value
	}
}

func WithLogFileMaxBackups(value int) applyOption {
	return func(o *options) {
		o.logFileMaxBackups = value
	}
}

func WithLogFileMaxAge(value int) applyOption {
	return func(o *options) {
		o.logFileMaxAge = value
	}
}

func WithTimeLocation(loc *time.Location) applyOption {
	return func(o *options) {
		o.timeLocation = loc
	}
}

func (options *options) getLogFilePath() string {
	return fmt.Sprintf("%s/%s", options.logDir, options.logFilename)
}

func (options *options) configLumperjack() (outputPaths, errorOutputPaths []string) {
	var ll = &lumberjack.Logger{
		Filename:   options.getLogFilePath(),
		MaxSize:    options.logFileMaxSize,
		MaxBackups: options.logFileMaxBackups,
		MaxAge:     options.logFileMaxAge,
		Compress:   options.logFileCompress,
	}

	zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: ll,
		}, nil
	})

	outputPaths = []string{
		fmt.Sprintf("lumberjack:%s", options.getLogFilePath()),
	}
	errorOutputPaths = []string{
		fmt.Sprintf("lumberjack:%s", options.getLogFilePath()),
	}

	return
}

func (options *options) ensureFile() error {
	var err = os.MkdirAll(options.logDir, 0755)
	return err

}
func (options *options) configDailyRotate() (outputPaths, errorOutputPaths []string, err error) {
	options.ensureFile()

	var dailyrotate *File
	var filePath = options.getLogFilePath()

	dailyrotate, err = NewFileWithPathGenerator(func(t time.Time) string {
		return filePath
	}, func(path string, didRotate bool) {

	})
	if err != nil {
		return
	}
	dailyrotate.Location = options.timeLocation
	zap.RegisterSink("dailyrotate", func(*url.URL) (zap.Sink, error) {
		return dailyrotateSink{
			dailyrotate,
		}, nil
	})

	outputPaths = []string{
		fmt.Sprintf("dailyrotate:%s", filePath),
	}
	errorOutputPaths = []string{
		fmt.Sprintf("dailyrotate:%s", filePath),
	}

	return
}
