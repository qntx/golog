package main

import (
	"log"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	gologzerolog "golog.qntx.fun/zerolog"
)

func main() {
	// Configure the lumberjack logger.
	// This will write to a file named "rotated_app.log" in the current directory.
	// It will rotate the log file when it reaches 1 megabyte.
	// It will keep up to 3 old log files.
	// Old log files will be kept for up to 7 days.
	// Compressed (zipped) backups are disabled by default, but can be enabled.
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log", // Log file name
		MaxSize:    1,         // megabytes
		MaxBackups: 3,         // number of old log files to keep
		MaxAge:     7,         // days
		Compress:   false,     // disabled by default
	}

	// Create a zerolog.Logger instance that writes to the lumberjack.Logger.
	// We want JSON output in the rotated files.
	// Set the timestamp and caller info.
	// Use the golog.CallerSkipFrameCount for correct caller info.
	rotatingFileLoggerInstance := zerolog.New(lumberjackLogger).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		CallerWithSkipFrameCount(gologzerolog.CallerSkipFrameCount).
		Logger()

	// Create a golog.Logger using NewWith.
	logger := gologzerolog.NewWith(rotatingFileLoggerInstance)

	// Log a bunch of messages to potentially trigger rotation.
	// In a real application, these would be actual log events.
	log.Println("Starting to log heavily to 'app.log' to demonstrate rotation...")
	for i := range 20000 { // Adjust count based on desired log volume
		logger.Infof("Logging message number %d. This is a somewhat long line to help fill up the log file relatively quickly for demonstration purposes.", i)
		if i%1000 == 0 {
			logger.Warnf("Milestone: Logged %d messages.", i)
		}
	}

	logger.Error("Finished logging a large number of messages.")
	log.Printf("Log messages written to 'app.log'. Check the file and potential rotated backups (e.g., app-YYYY-MM-DDTHH-MM-SS.mmm.log).")
}
