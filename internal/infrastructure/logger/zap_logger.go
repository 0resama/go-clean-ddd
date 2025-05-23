package logger

import "go.uber.org/zap"

type ZapLogger struct {
    logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
    logger, _ := zap.NewProduction()
    return &ZapLogger{logger: logger}
}

func (l *ZapLogger) Info(msg string, fields ...any) {
    l.logger.Sugar().Infow(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...any) {
    l.logger.Sugar().Errorw(msg, fields...)
}