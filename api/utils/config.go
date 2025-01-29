package utils

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Config holds all the environment variables.
type Config struct {
	SecretKey      string
	SurfReportLink string
	LogLevel       string
}

// Initialize a global SugaredLogger
var SugaredLogger *zap.SugaredLogger

func init() {
	// Initialize the logger
	logger, _ := zap.NewProduction()
	SugaredLogger = logger.Sugar()
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		SugaredLogger.Warnw("Warning: .env file not found",
			"error", err,
		)
	}

	return &Config{
		SecretKey:      getEnv("SECRET_KEY", "your_secret_key"),
		SurfReportLink: getEnv("SURF_REPORT_LINK", "https://www.surf-report.com/meteo-surf/la-guerite-s1170.html"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		SugaredLogger.Debugw("Using default value for config",
			"key", key,
			"default", defaultValue,
		)
		return defaultValue
	}
	return value
}
