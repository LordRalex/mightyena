package logging

type Logger interface {
	Log(level Level, format string, args ...interface{})
}

var cache = make(map[string]Logger)

func GetLogger(name string) Logger {

	logger := cache[name]
	if logger == nil {
		logger = &prefixLogger{Name: name}
		cache[name] = logger
	}

	return logger
}

