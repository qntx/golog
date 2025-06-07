# Golog

A minimalist logging tool built on `zerolog`. Less is more.

## Features

* Levels: `Debug`, `Info`, `Warn`, `Error`, `Fatal`
* Lightweight, efficient, extensible

## Usage

### 1. Basic - Log to Console

```go
logger := logger.New("info", os.Stdout)
logger.Info("Hello, world!")
```

### 2. Log to File

```go
file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
logger := logger.New("debug", file)
logger.Debug("Saved to file")
```

### 3. Log to Console and File

```go
file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
mw := io.MultiWriter(os.Stdout, file)
logger := logger.New("info", mw)
logger.Info("Dual output")
```

### 4. Log Rotation

```go
rotator := &lumberjack.Logger{
    Filename:   "app.log",
    MaxSize:    10, // MB
    MaxBackups: 3,
    MaxAge:     28, // days
}
logger := logger.New("warn", rotator)
logger.Warn("Rotation enabled")
```

> *Note: Requires lumberjack package.*
>
> ```bash
> go get github.com/natefinch/lumberjack
> ```

### 5. Log Rotation + Console

```go
rotator := &lumberjack.Logger{
    Filename:   "app.log",
    MaxSize:    500, // MB
    MaxBackups: 3,
    MaxAge:     28, // days
    Compress:   true,
}
mw := io.MultiWriter(os.Stdout, rotator)
logger := logger.New("info", mw)
logger.Info("Rotation and console")
```

### 6. Async Logging

```go
file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
logger := logger.New("error", file)
go logger.Error("Async logging")
```

## Get Started

```go
import "github.com/Qntx/qutil/logger"
```

Simple yet powerful—logging, refined.
