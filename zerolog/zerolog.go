// Package zerolog provides an implementation of golog.Logger using rs/zerolog.
package zerolog

import (
	"os"

	"github.com/rs/zerolog"

	"golog.qntx.fun"
)

// CallerSkipFrameCount is the number of stack frames to skip to find the caller.
// It is set to zerolog.CallerSkipFrameCount + 1 to account for the golog wrapper.
var CallerSkipFrameCount = zerolog.CallerSkipFrameCount + 1

// logger wraps a zerolog.Logger with the golog.Logger interface.
//
// Provides leveled logging with configurable timestamp and caller details.
// Example:
//
//	l := New()
//	l.Info("Server started on port 8080")
type logger struct {
	l *zerolog.Logger // Embedded zerolog.Logger for efficiency.
}

var _ golog.Logger = (*logger)(nil)

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
func New() golog.Logger {
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
func NewWith(log zerolog.Logger) golog.Logger {
	return &logger{l: &log}
}

// Level dynamically updates the logging level.
func (l *logger) Level(level golog.Level) {
	*l.l = l.l.Level(zerolog.Level(level))
}

// GetLevel returns the current logging level.
func (l *logger) GetLevel() golog.Level {
	return golog.Level(l.l.GetLevel())
}

// Trace logs a message at the trace level.
func (l *logger) Trace(msg string) {
	l.l.Trace().Msg(msg)
}

// Tracef logs a formatted message at the trace level.
func (l *logger) Tracef(format string, args ...any) {
	l.l.Trace().Msgf(format, args...)
}

// Debug logs a message at the debug level.
func (l *logger) Debug(msg string) {
	l.l.Debug().Msg(msg)
}

// Debugf logs a formatted message at the debug level.
func (l *logger) Debugf(format string, args ...any) {
	l.l.Debug().Msgf(format, args...)
}

// Info logs a message at the info level.
func (l *logger) Info(msg string) {
	l.l.Info().Msg(msg)
}

// Infof logs a formatted message at the info level.
func (l *logger) Infof(format string, args ...any) {
	l.l.Info().Msgf(format, args...)
}

// Warn logs a message at the warn level.
func (l *logger) Warn(msg string) {
	l.l.Warn().Msg(msg)
}

// Warnf logs a formatted message at the warn level.
func (l *logger) Warnf(format string, args ...any) {
	l.l.Warn().Msgf(format, args...)
}

// Error logs a message at the error level.
func (l *logger) Error(msg string) {
	l.l.Error().Msg(msg)
}

// Errorf logs a formatted message at the error level.
func (l *logger) Errorf(format string, args ...any) {
	l.l.Error().Msgf(format, args...)
}

// Fatal logs a message at the fatal level and exits with os.Exit(1).
func (l *logger) Fatal(msg string) {
	l.l.Fatal().Msg(msg)
}

// Fatalf logs a formatted message at the fatal level and exits with os.Exit(1).
func (l *logger) Fatalf(format string, args ...any) {
	l.l.Fatal().Msgf(format, args...)
}

// Panic logs a message at the panic level and calls panic().
func (l *logger) Panic(msg string) {
	l.l.Panic().Msg(msg)
}

// Panicf logs a formatted message at the panic level and calls panic().
func (l *logger) Panicf(format string, args ...any) {
	l.l.Panic().Msgf(format, args...)
}
