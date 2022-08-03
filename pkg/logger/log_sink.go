package logger

import "gopkg.in/natefinch/lumberjack.v2"

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

type dailyrotateSink struct {
	*File
}

func (dailyrotateSink) Sync() error {
	return nil
}
