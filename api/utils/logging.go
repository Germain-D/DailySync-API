package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

// Initialize sets up the Zap logger with the specified log level.
func Initialize(logLevel string) error {
	// Parse the log level
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	// Configure the logger
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(config)
	atomicLevel := zap.NewAtomicLevelAt(level)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), atomicLevel)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Initialize SugaredLogger
	SugaredLogger = Logger.Sugar()

	// Replace the global logger
	zap.ReplaceGlobals(Logger)

	return nil
}

// Sync flushes any buffered log entries.
func Sync() {
	_ = Logger.Sync()
	_ = SugaredLogger.Sync()
}
