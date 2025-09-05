package golog

import "github.com/rs/zerolog"

// Level represents a log severity level. Use the package variables as an enum.
type Level zerolog.Level

const (
	LevelTrace   = Level(zerolog.TraceLevel)
	LevelDebug   = Level(zerolog.DebugLevel)
	LevelInfo    = Level(zerolog.InfoLevel)
	LevelWarn    = Level(zerolog.WarnLevel)
	LevelError   = Level(zerolog.ErrorLevel)
	LevelFatal   = Level(zerolog.FatalLevel)
	LevelPanic   = Level(zerolog.PanicLevel)
	LevelNoLevel = Level(zerolog.NoLevel)
)

func (l Level) String() string {
	return zerolog.Level(l).String()
}

// ParseLevel parses a string-based level and returns the corresponding Level.
//
// Supported strings are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL, and
// their lower-case forms. Returns an error if the level is invalid.
func ParseLevel(level string) (Level, error) {
	l, err := zerolog.ParseLevel(level)
	return Level(l), err
}

// Logger defines the contract for logging operations.
type Logger interface {
	Trace(msg string)
	Tracef(format string, args ...any)
	Debug(msg string)
	Debugf(format string, args ...any)
	Info(msg string)
	Infof(format string, args ...any)
	Warn(msg string)
	Warnf(format string, args ...any)
	Error(msg string)
	Errorf(format string, args ...any)
	Fatal(msg string)
	Fatalf(format string, args ...any)
	Panic(msg string)
	Panicf(format string, args ...any)
	Level(level Level)
	GetLevel() Level
}
