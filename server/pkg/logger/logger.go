package logger

import "go.uber.org/zap"

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
}

type logger struct {
	logger *zap.Logger
}

func NewLogger(isDevMode bool) (Logger, error) {
	var (
		l   *zap.Logger
		err error
	)
	if isDevMode {
		l, err = zap.NewProduction()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}

	return &logger{logger: l}, nil
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}
