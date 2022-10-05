package logger

import (
	"os"
	"time"

	"github.com/ArifulProtik/chatserver-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

type applogger struct {
	log *zap.SugaredLogger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 01, 2006  15:04:05"))
}

func New(cfg *config.App) *applogger {
	level := loggerLevelMap["debug"]
	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if cfg.Status == "Dev" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	encoderCfg.EncodeTime = SyslogTimeEncoder
	encoder = zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(level)) // generic log level
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &applogger{
		log: logger.Sugar(),
	}

}
func (l *applogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *applogger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args...)
}
func (l *applogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *applogger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args...)
}
