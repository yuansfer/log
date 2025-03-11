package log

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestInitLog(t *testing.T) {
	err := InitLog()
	if err != nil {
		t.Errorf("InitLog() error = %v", err)
	}
}

func TestInfo(t *testing.T) {
	InitLog()
	Info("This is an info message")
}

func TestWarn(t *testing.T) {
	InitLog()
	Warn("This is a warning message")
}

func TestError(t *testing.T) {
	InitLog()
	Error("This is an error message")
}

func TestDebug(t *testing.T) {
	InitLog()
	Debug("This is a debug message")
}

func TestInfof(t *testing.T) {
	InitLog()
	Infof("This is an info message with %s", "formatting")
}

func TestWarnf(t *testing.T) {
	InitLog()
	Warnf("This is a warning message with %s", "formatting")
}

func TestErrorf(t *testing.T) {
	InitLog()
	Errorf("This is an error message with %s", "formatting")
}

func TestDebugf(t *testing.T) {
	InitLog()
	Debugf("This is a debug message with %s", "formatting")
}

func TestSetLoggerLevel(t *testing.T) {
	InitLog()
	yLogger, ok := logger.(*YLogger)
	if !ok {
		t.Fatalf("logger is not of type *YLogger")
	}
	yLogger.SetLoggerLevel("debug")
	if yLogger.dynamicLevel.Level() != zapcore.DebugLevel {
		t.Errorf("expected debug level, got %v", yLogger.dynamicLevel.Level())
	}
}
