// Package golog provides a flexible, high-performance logging system using zerolog.
//
// It supports leveled logging (debug, info, warn, error, fatal, panic) with timestamps and
// caller information, designed for production-grade applications. Not thread-safe
// unless the underlying io.Writer is.
package golog

import (
	"os"

	"github.com/rs/zerolog"
)

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

// CallerSkipFrameCount is the number of stack frames to skip to find the caller.
// It is set to zerolog.CallerSkipFrameCount + 1 to account for the golog wrapper.
var CallerSkipFrameCount = zerolog.CallerSkipFrameCount + 1

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

// Interface defines the contract for logging operations.
type Interface interface {
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

// Logger wraps a zerolog.Logger with a custom interface.
//
// Provides leveled logging with configurable timestamp and caller details.
// Example:
//
//	l := New()
//	l.Info("Server started on port 8080")
type Logger struct {
	l *zerolog.Logger // Embedded zerolog.Logger for efficiency.
}

var _ Interface = (*Logger)(nil)

// New creates a new Logger with default configuration.
//
// Defaults:
//   - Level: LevelInfo
//   - Writer: os.Stdout
//   - Timestamps: Enabled
//   - Caller Info: Enabled
//
// Example:
//
//	l := New()
//	l.Info("Starting application")
func New() *Logger {
	return NewWith(
		zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			CallerWithSkipFrameCount(CallerSkipFrameCount).
			Logger())
}

// NewWith creates a new Logger with the specified configuration.
//
// Example:
//
//	l := NewWith(zerolog.New(os.Stdout))
//	l.Debug("Debugging enabled")
func NewWith(log zerolog.Logger) *Logger {
	return &Logger{l: &log}
}

// Level dynamically updates the logging level.
func (l *Logger) Level(level Level) {
	*l.l = l.l.Level(zerolog.Level(level))
}

// GetLevel returns the current logging level.
func (l *Logger) GetLevel() Level {
	return Level(l.l.GetLevel())
}

// Trace logs a message at the trace level.
func (l *Logger) Trace(msg string) {
	l.l.Trace().Msg(msg)
}

// Tracef logs a formatted message at the trace level.
func (l *Logger) Tracef(format string, args ...any) {
	l.l.Trace().Msgf(format, args...)
}

// Debug logs a message at the debug level.
func (l *Logger) Debug(msg string) {
	l.l.Debug().Msg(msg)
}

// Debugf logs a formatted message at the debug level.
func (l *Logger) Debugf(format string, args ...any) {
	l.l.Debug().Msgf(format, args...)
}

// Info logs a message at the info level.
func (l *Logger) Info(msg string) {
	l.l.Info().Msg(msg)
}

// Infof logs a formatted message at the info level.
func (l *Logger) Infof(format string, args ...any) {
	l.l.Info().Msgf(format, args...)
}

// Warn logs a message at the warn level.
func (l *Logger) Warn(msg string) {
	l.l.Warn().Msg(msg)
}

// Warnf logs a formatted message at the warn level.
func (l *Logger) Warnf(format string, args ...any) {
	l.l.Warn().Msgf(format, args...)
}

// Error logs a message at the error level.
func (l *Logger) Error(msg string) {
	l.l.Error().Msg(msg)
}

// Errorf logs a formatted message at the error level.
func (l *Logger) Errorf(format string, args ...any) {
	l.l.Error().Msgf(format, args...)
}

// Fatal logs a message at the fatal level and exits with os.Exit(1).
func (l *Logger) Fatal(msg string) {
	l.l.Fatal().Msg(msg)
}

// Fatalf logs a formatted message at the fatal level and exits with os.Exit(1).
func (l *Logger) Fatalf(format string, args ...any) {
	l.l.Fatal().Msgf(format, args...)
}

// Panic logs a message at the panic level and calls panic().
func (l *Logger) Panic(msg string) {
	l.l.Panic().Msg(msg)
}

// Panicf logs a formatted message at the panic level and calls panic().
func (l *Logger) Panicf(format string, args ...any) {
	l.l.Panic().Msgf(format, args...)
}
