package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	logger Logger
)

type YLogger struct {
	Logger
	dynamicLevel zap.AtomicLevel
}

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
}

func init() {
	if err := InitLog(); nil != err {
		log.Println(err.Error())
	}
}

func InitLog() error {
	conf := &zap.Config{
		Level:             zap.NewAtomicLevel(),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     nil,
			EncodeLevel:    nil,
		},
	}
	conf.OutputPaths = []string{`stderr`}
	conf.ErrorOutputPaths = []string{`stderr`}
	InitLogger(conf)

	return nil
}

func InitLogger(conf *zap.Config) {
	var zapLoggerConfig zap.Config
	if conf == nil {
		zapLoggerConfig = zap.NewDevelopmentConfig()
		zapLoggerEncoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		zapLoggerConfig.EncoderConfig = zapLoggerEncoderConfig
	} else {
		zapLoggerConfig = *conf
	}
	zapLogger, _ := zapLoggerConfig.Build(zap.AddCallerSkip(1))
	logger = &YLogger{Logger: zapLogger.Sugar(), dynamicLevel: zapLoggerConfig.Level}
}

type OpsLogger interface {
	Logger
	SetLoggerLevel(level string)
}

func (dl *YLogger) SetLoggerLevel(level string) {
	l := new(zapcore.Level)
	l.Set(level)
	dl.dynamicLevel.SetLevel(*l)
}
