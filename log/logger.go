package log

import (
	"fmt"
	"io"

	"github.com/Sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

func (l *Logger) SetOutput(out io.Writer) {
	l.logger.Out = out
}

func (l *Logger) SetJSONFormatter() {
	l.logger.Formatter = new(logrus.JSONFormatter)
}

func (l *Logger) Output() io.Writer {
	return l.logger.Out
}

func (l *Logger) SetLevel(level string) {
	switch level {
	case "debug":
		l.logger.Level = logrus.DebugLevel
	case "info":
		l.logger.Level = logrus.InfoLevel
	case "warn":
		l.logger.Level = logrus.WarnLevel
	case "error":
		l.logger.Level = logrus.ErrorLevel
	case "fatal":
		l.logger.Level = logrus.FatalLevel
	case "panic":
		l.logger.Level = logrus.PanicLevel
	default:
		panic(fmt.Errorf("%s: config log.level `%s` invalid.", pkgName, level))
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

func (l *Logger) Warningln(args ...interface{}) {
	l.logger.Warningln(args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}
