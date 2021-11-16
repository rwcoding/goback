package logger

func InitDefaultLogger() {
	newZap()
}

func Debug(template string, args ...interface{}) {
	sugar.Debugf(template, args)
}

func Info(template string, args ...interface{}) {
	sugar.Infof(template, args)
}

func Warn(template string, args ...interface{}) {
	sugar.Warnf(template, args)
}

func Error(template string, args ...interface{}) {
	sugar.Errorf(template, args)
}

func Fatal(template string, args ...interface{}) {
	sugar.Fatalf(template, args)
}

func Panic(template string, args ...interface{}) {
	sugar.Panicf(template, args)
}
