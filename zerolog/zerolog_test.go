package zerolog_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/qntx/golog"
	gologzerolog "github.com/qntx/golog/zerolog"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultLevel(t *testing.T) {
	l := gologzerolog.New()
	assert.Equal(t, golog.LevelInfo, l.GetLevel())
}

func TestNewWithCustomLogger(t *testing.T) {
	zl := zerolog.New(os.Stderr).Level(zerolog.DebugLevel)
	l := gologzerolog.NewWith(zl)
	assert.Equal(t, golog.LevelDebug, l.GetLevel())
}

func TestLevelSetAndGet(t *testing.T) {
	l := gologzerolog.New()
	l.Level(golog.LevelDebug)
	assert.Equal(t, golog.LevelDebug, l.GetLevel())
}

func TestTraceMsg(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.TraceLevel)
	l := gologzerolog.NewWith(zl)
	l.Trace("trace message")
	assert.Contains(t, buf.String(), "trace message")
}

func TestTracef(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.TraceLevel)
	l := gologzerolog.NewWith(zl)
	l.Tracef("trace %s", "formatted")
	assert.Contains(t, buf.String(), "trace formatted")
}

func TestDebugMsg(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.DebugLevel)
	l := gologzerolog.NewWith(zl)
	l.Debug("debug message")
	assert.Contains(t, buf.String(), "debug message")
}

func TestDebugf(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.DebugLevel)
	l := gologzerolog.NewWith(zl)
	l.Debugf("debug %s", "formatted")
	assert.Contains(t, buf.String(), "debug formatted")
}

func TestInfoMsg(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.InfoLevel)
	l := gologzerolog.NewWith(zl)
	l.Info("info message")
	assert.Contains(t, buf.String(), "info message")
}

func TestInfof(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.InfoLevel)
	l := gologzerolog.NewWith(zl)
	l.Infof("info %s", "formatted")
	assert.Contains(t, buf.String(), "info formatted")
}

func TestWarnMsg(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.WarnLevel)
	l := gologzerolog.NewWith(zl)
	l.Warn("warn message")
	assert.Contains(t, buf.String(), "warn message")
}

func TestWarnf(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.WarnLevel)
	l := gologzerolog.NewWith(zl)
	l.Warnf("warn %s", "formatted")
	assert.Contains(t, buf.String(), "warn formatted")
}

func TestErrorMsg(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.ErrorLevel)
	l := gologzerolog.NewWith(zl)
	l.Error("error message")
	assert.Contains(t, buf.String(), "error message")
}

func TestErrorf(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.ErrorLevel)
	l := gologzerolog.NewWith(zl)
	l.Errorf("error %s", "formatted")
	assert.Contains(t, buf.String(), "error formatted")
}

func TestPanicMsg(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.PanicLevel)
	l := gologzerolog.NewWith(zl)
	assert.PanicsWithValue(t, "panic message", func() {
		l.Panic("panic message")
	})
	assert.Contains(t, buf.String(), "panic message")
}

func TestPanicf(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf).Level(zerolog.PanicLevel)
	l := gologzerolog.NewWith(zl)
	assert.PanicsWithValue(t, "panic formatted", func() {
		l.Panicf("panic %s", "formatted")
	})
	assert.Contains(t, buf.String(), "panic formatted")
}
