package database

import (
	"context"
	"errors"
	"github.com/cesc1802/auth-service/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var quoteEscaper = strings.NewReplacer("\"", "", "\t", "", "\n", " ")

func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type Logger struct {
	ZapLogger                 *logger.Logger
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func DefaultLogger(namspace string) *Logger {
	var dbLogger = &Logger{
		ZapLogger:                 logger.New(namspace),
		LogLevel:                  gormLogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
	gormLogger.Default = dbLogger
	return dbLogger
}
func (l Logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return Logger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormLogger.Info {
		return
	}
	l.logger().Sugar().Debugf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormLogger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormLogger.Error {
		return
	}
	l.logger().Sugar().Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	var elapsed = time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormLogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		sql = EscapeQuotes(sql)
		l.logger().Error("trace", zap.Error(err), zap.Duration("sql_elapsed", elapsed), zap.Int64("sql_rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		sql = EscapeQuotes(sql)
		l.logger().Warn("trace", zap.Duration("sql_elapsed", elapsed), zap.Int64("sql_rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gormLogger.Info:
		sql, rows := fc()
		sql = EscapeQuotes(sql)
		l.logger().Debug("trace", zap.Duration("sql_elapsed", elapsed), zap.Int64("sql_rows", rows), zap.String("sql", sql))
	}
}

var (
	gormPackage  = filepath.Join("gorm.io", "gorm")
	queryPackage = filepath.Join("/app/pkg/repo/query/queryfunc", "queryfunc")
	dbPackage    = filepath.Join("/app/pkg", "db")
)

func (l Logger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, queryPackage):
		case strings.Contains(file, dbPackage):
		default:
			return l.ZapLogger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger.Logger
}
