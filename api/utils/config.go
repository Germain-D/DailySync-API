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
	StormGlassKey  string
	SpotSurfLat    string
	SpotSurfLon    string
	WeatherLat     string
	WeatherLon     string
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
		StormGlassKey:  getEnv("STORM_GLASS_KEY", "your_stormglass_key"),
		SpotSurfLat:    getEnv("SPOT_SURF_LAT", "47.593"),
		SpotSurfLon:    getEnv("SPOT_SURF_LON", "-3.148"),
		WeatherLat:     getEnv("WEATHER_LAT", "47.6569"),
		WeatherLon:     getEnv("WEATHER_LON", "-2.762"),
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
