package log

import (
	"io"
)

var (
	_std = New()
)

func SetOutput(out io.Writer) {
	_std.SetOutput(out)
}

func SetJSONFormatter() {
	_std.SetJSONFormatter()
}

func Output() io.Writer {
	return _std.Output()
}

func SetLevel(level string) {
	_std.SetLevel(level)
}

// exported std api
func Debugf(format string, args ...interface{})   { _std.Debugf(format, args...) }
func Infof(format string, args ...interface{})    { _std.Infof(format, args...) }
func Printf(format string, args ...interface{})   { _std.Printf(format, args...) }
func Warnf(format string, args ...interface{})    { _std.Warnf(format, args...) }
func Warningf(format string, args ...interface{}) { _std.Warningf(format, args...) }
func Errorf(format string, args ...interface{})   { _std.Errorf(format, args...) }
func Fatalf(format string, args ...interface{})   { _std.Fatalf(format, args...) }
func Panicf(format string, args ...interface{})   { _std.Panicf(format, args...) }
func Debug(args ...interface{})                   { _std.Debug(args...) }
func Info(args ...interface{})                    { _std.Info(args...) }
func Print(args ...interface{})                   { _std.Print(args...) }
func Warn(args ...interface{})                    { _std.Warn(args...) }
func Warning(args ...interface{})                 { _std.Warning(args...) }
func Error(args ...interface{})                   { _std.Error(args...) }
func Fatal(args ...interface{})                   { _std.Fatal(args...) }
func Panic(args ...interface{})                   { _std.Panic(args...) }
func Debugln(args ...interface{})                 { _std.Debugln(args...) }
func Infoln(args ...interface{})                  { _std.Infoln(args...) }
func Println(args ...interface{})                 { _std.Println(args...) }
func Warnln(args ...interface{})                  { _std.Warnln(args...) }
func Warningln(args ...interface{})               { _std.Warningln(args...) }
func Errorln(args ...interface{})                 { _std.Errorln(args...) }
func Fatalln(args ...interface{})                 { _std.Fatalln(args...) }
func Panicln(args ...interface{})                 { _std.Panicln(args...) }
