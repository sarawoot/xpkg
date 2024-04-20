package xlog

var logger *Logger

func init() {
	logger = new()
}

func Info(msg string) {
	logger.Info(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Fatal(msg string) {
	logger.Fatal(msg)
}

func WithFields(kv map[string]interface{}) *Logger {
	l := new().logger.With().Fields(kv).Logger()
	return &Logger{logger: l}
}
