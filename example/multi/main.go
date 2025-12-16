package main

import (
	"io"
	"log"
	"os"

	"github.com/rs/zerolog"
	gologzerolog "golog.qntx.fun/zerolog"
)

func main() {
	// --- Setup for File Output ---
	const logFileName = "app.log"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", logFileName, err)
	}
	defer file.Close()

	// --- Setup for Console Output (Human-readable) ---
	// We'll use ConsoleWriter for pretty printing to os.Stderr
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}

	// --- Create MultiWriter ---
	// This will write to both the console (via ConsoleWriter) and the file.
	// Note: The file will receive JSON output by default from zerolog.New().
	// If you want console-like output in the file as well, you'd need another ConsoleWriter for the file.
	// For this example, file gets JSON, console gets pretty print.
	multiWriter := io.MultiWriter(consoleWriter, file)

	// Create a zerolog.Logger instance that writes to the multiWriter.
	// We set a base level here, for example, Debug.
	// Timestamp and Caller will be included.
	multiLoggerInstance := zerolog.New(multiWriter).
		Level(zerolog.DebugLevel). // Set a base level for the zerolog instance
		With().
		Timestamp().
		CallerWithSkipFrameCount(gologzerolog.CallerSkipFrameCount). // Caller will be based on where golog's methods are called
		Logger()

	// Create a golog.Logger using NewWith.
	logger := gologzerolog.NewWith(multiLoggerInstance)

	// Log some messages
	logger.Info("This informational message goes to both console and file.")
	logger.Warnf("This warning message (formatted: %s) also goes to both.", "multi-output")
	logger.Debug("A debug message for console and file.")
	logger.Error("An error message for all outputs.")

	log.Printf("Log messages written to console and %s. Check the file content.", logFileName)
}
