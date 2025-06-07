package main

import (
	"github.com/qntx/golog"
)

func main() {
	// Create a new logger with default settings.
	// It will log Info level and above to os.Stdout (via os.Stderr with ConsoleWriter)
	// and include timestamps and caller information.
	logger := golog.New()

	logger.Debug("This debug message will not be printed by default.") // Default level is Info
	logger.Trace("Hello, world from the basic console example!")       // Trace level is not printed by default
	logger.Infof("Hello %s from the basic console example!", "world")

	// Example of changing the level
	logger.Level(golog.LevelTrace)
	logger.Trace("This trace message WILL be printed after changing the level.")
	logger.Debug("This debug message WILL be printed after changing the level.")
	logger.Errorf("Hello %s from the basic console example!", "world")
}
