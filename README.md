# Go log

A minimalist, high-performance logging library built on [`zerolog`](https://github.com/rs/zerolog).

## Features

- Leveled logging: Trace, Debug, Info, Warn, Error, Fatal, Panic
- Zero-allocation JSON logging
- Configurable output: console, file, multi-writer
- Dynamic level changes, timestamps, caller info

## Installation

```bash
go get github.com/qntx/golog
```

## Usage

See the [`example/`](./example) folder for usage examples, including:

- Basic console logging
- File logging
- Console and file output
- Log rotation with `lumberjack`
- Dynamic log level changes

## Notes

- **Thread Safety**: Not thread-safe unless the underlying `io.Writer` is.
- **Async Logging**: Use goroutines with proper synchronization.
- **Customization**: Extend with zerolog features (e.g., custom fields, hooks).

## License

MIT
