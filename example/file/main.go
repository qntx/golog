package main

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"golog.qntx.fun"
	gologzerolog "golog.qntx.fun/zerolog"
)

func main() {
	// Define the log file name
	const logFileName = "app.log"

	// Open the log file for writing.
	// os.O_CREATE: create the file if it doesn't exist.
	// os.O_WRONLY: open the file write-only.
	// os.O_APPEND: append to the end of the file.
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", logFileName, err)
	}
	// It's good practice to close the file when main exits.
	defer file.Close()

	// Create a zerolog.Logger instance that writes to the file.
	// We are not using ConsoleWriter here because we want JSON output in the file.
	fileLoggerInstance := zerolog.New(file).With().Timestamp().CallerWithSkipFrameCount(gologzerolog.CallerSkipFrameCount).Logger()

	// Create a golog.Logger using NewWith.
	logger := gologzerolog.NewWith(fileLoggerInstance)

	// Log some messages
	logger.Info("This is an informational message logged to the file.")
	logger.Warnf("This is a warning message with formatting: %s", "something important")
	logger.Error("This is an error message.")

	// You can also change the level for the file logger if needed
	logger.Level(golog.LevelDebug)
	logger.Debug("This debug message will also be logged to the file.")
	logger.Errorf("Hello %s from the basic console example!", "world")

	log.Printf("Log messages written to %s. Check its content.", logFileName)
}
