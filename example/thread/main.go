package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	gologzerolog "github.com/qntx/golog/zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

func main() {
	// Configure the diode.Writer
	// - os.Stdout: The actual underlying writer (can be a file, os.Stderr, etc.)
	// - 1000:      The size of the buffer (number of log events).
	// - 10 * time.Millisecond: The polling interval for the writer goroutine.
	// - func(missed int): Callback function invoked when messages are dropped.
	diodeWriter := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		// This message will be printed to standard fmt output, not through the logger,
		// to avoid potential deadlocks or further drops if the logger itself is overwhelmed.
		fmt.Printf("Logger Dropped %d messages\n", missed)
	})

	// Create a zerolog.Logger instance that writes to the diodeWriter.
	// We'll use ConsoleWriter for prettier output in this example,
	// but you could use plain JSON output too.
	// The diodeWriter itself handles the async and non-blocking aspects.
	consoleDiodeWriter := zerolog.ConsoleWriter{Out: diodeWriter, TimeFormat: time.Kitchen}

	asyncLoggerInstance := zerolog.New(consoleDiodeWriter).
		Level(zerolog.DebugLevel). // Set a base level
		With().
		Timestamp().
		CallerWithSkipFrameCount(gologzerolog.CallerSkipFrameCount). // Ensure correct caller info
		Logger()

	// Create a golog.Logger using NewWith.
	logger := gologzerolog.NewWith(asyncLoggerInstance)

	logger.Info("Starting application with thread-safe, non-blocking logger.")

	// Simulate concurrent logging from multiple goroutines
	var wg sync.WaitGroup
	numGoroutines := 50
	logsPerGoroutine := 200 // Increase this to potentially see dropped messages if buffer is small or writer is slow

	for i := range numGoroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range logsPerGoroutine {
				logger.Infof("Log from goroutine %d, message %d", id, j)
				// time.Sleep(1 * time.Microsecond) // Optionally slow down producers
			}
		}(i)
	}

	wg.Wait()
	logger.Info("All goroutines finished logging.")

	// Important: To ensure all buffered messages are flushed from the diode,
	// especially if the underlying writer (os.Stdout here) is buffered or if the
	// poller hasn't had a chance to run, you might need to give it a moment
	// or, for file writers, explicitly close the diodeWriter if it implements io.Closer
	// (diode.Writer itself does not directly, but the underlying writer might).
	// For os.Stdout, a small sleep is often enough for demonstration.
	// In a real application, you'd manage the lifecycle of the diodeWriter more carefully.
	time.Sleep(100 * time.Millisecond) // Give time for poller to flush
	logger.Info("Application finished.")

	// Note: diode.Writer does not implement io.Closer.
	// If the underlying writer (e.g., an os.File) needs to be closed,
	// you must manage its lifecycle separately.
	// For this example with os.Stdout, no explicit close is needed for the diode itself.
}
