package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const developmentEnvironment = "development"

type Config struct {
	Environment    string
	Host           string
	Port           string
	Database       DatabaseConfig
	Redis          RedisConfig
	Security       SecurityConfig
	Firebase       FirebaseConfig
	SermonCloud    SermonCloudConfig
	Stripe         StripeConfig
	PlanningCenter PlanningCenterConfig
	Email          EmailConfig
	Telemetry      TelemetryConfig
	Uptime         UptimeConfig
}

type DatabaseConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

type SecurityConfig struct {
	APIKey string
}

type FirebaseConfig struct {
	ProjectID string
}

type SermonCloudConfig struct {
	URL string
}

type StripeConfig struct {
	SecretKey      string
	PublishableKey string
	Version        string
}

type PlanningCenterConfig struct {
	URL      string
	TokenURL string
}

type EmailConfig struct {
	AppPassword string
}

type TelemetryConfig struct {
	ServiceName  string
	OTLPEndpoint string
}

type UptimeConfig struct {
	RequireBearer   bool
	ServiceID       string
	ServiceName     string
	KeyPrefix       string
	RetentionDays   int
	SampleIntervalS int
}

type Lookup func(string) (string, bool)

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && os.Getenv("ENVIRONMENT") == developmentEnvironment {
		return nil, fmt.Errorf("load .env file: %w", err)
	}

	return LoadConfigFrom(os.LookupEnv)
}

func LoadConfigFrom(lookup Lookup) (*Config, error) {
	uptimeRetentionDays, err := lookupInt(lookup, "FIBER_UPTIME_RETENTION_DAYS")
	if err != nil {
		return nil, err
	}
	uptimeSampleInterval, err := lookupInt(lookup, "FIBER_UPTIME_SAMPLE_INTERVAL_SECONDS")
	if err != nil {
		return nil, err
	}
	uptimeRequireBearer, err := lookupBoolWithDefault(lookup, "FIBER_UPTIME_REQUIRE_BEARER_AUTH", true)
	if err != nil {
		return nil, err
	}

	configuration := &Config{
		Environment: lookupValue(lookup, "ENVIRONMENT"),
		Host:        lookupValue(lookup, "HOST"),
		Port:        lookupValue(lookup, "PORT"),
		Database: DatabaseConfig{
			URL: lookupValue(lookup, "DATABASE_URL"),
		},
		Redis: RedisConfig{
			URL: lookupValue(lookup, "REDIS_URL"),
		},
		Security: SecurityConfig{
			APIKey: lookupValue(lookup, "API_KEY"),
		},
		Firebase: FirebaseConfig{
			ProjectID: lookupValue(lookup, "FIREBASE_PROJECT_ID"),
		},
		SermonCloud: SermonCloudConfig{
			URL: lookupValue(lookup, "SERMONCLOUD_URL"),
		},
		Stripe: StripeConfig{
			SecretKey:      lookupValue(lookup, "STRIPE_SECRET_KEY"),
			PublishableKey: lookupValue(lookup, "STRIPE_PUBLISHABLE_KEY"),
			Version:        lookupValue(lookup, "STRIPE_VERSION"),
		},
		PlanningCenter: PlanningCenterConfig{
			URL:      lookupValue(lookup, "PLANNING_CENTER_URL"),
			TokenURL: lookupValue(lookup, "PLANNING_CENTER_TOKEN_URL"),
		},
		Email: EmailConfig{
			AppPassword: lookupValue(lookup, "EMAIL_APP_PASSWORD"),
		},
		Telemetry: TelemetryConfig{
			ServiceName:  lookupValue(lookup, "OTEL_SERVICE_NAME"),
			OTLPEndpoint: lookupValue(lookup, "OTEL_EXPORTER_OTLP_ENDPOINT"),
		},
		Uptime: UptimeConfig{
			RequireBearer:   uptimeRequireBearer,
			ServiceID:       lookupValue(lookup, "FIBER_UPTIME_SERVICE_ID"),
			ServiceName:     lookupValue(lookup, "FIBER_UPTIME_SERVICE_NAME"),
			KeyPrefix:       lookupValue(lookup, "FIBER_UPTIME_KEY_PREFIX"),
			RetentionDays:   uptimeRetentionDays,
			SampleIntervalS: uptimeSampleInterval,
		},
	}

	if configuration.Environment == "" {
		configuration.Environment = developmentEnvironment
	}
	if configuration.Host == "" {
		configuration.Host = "127.0.0.1"
		if configuration.Environment == "production" {
			configuration.Host = "0.0.0.0"
		}
	}
	if configuration.Port == "" {
		configuration.Port = "8080"
	}
	if configuration.SermonCloud.URL == "" {
		configuration.SermonCloud.URL = "https://api.sermoncloud.com/momentum-church-1/sermons"
	}
	if configuration.Telemetry.ServiceName == "" {
		configuration.Telemetry.ServiceName = "momentum-api"
	}
	if configuration.Uptime.ServiceID == "" {
		configuration.Uptime.ServiceID = "momentum-api"
	}
	if configuration.Uptime.ServiceName == "" {
		configuration.Uptime.ServiceName = "Momentum API"
	}
	if configuration.Uptime.KeyPrefix == "" {
		configuration.Uptime.KeyPrefix = "momentum:uptime"
	}
	if configuration.Uptime.RetentionDays <= 0 {
		configuration.Uptime.RetentionDays = 30
	}
	if configuration.Uptime.SampleIntervalS <= 0 {
		configuration.Uptime.SampleIntervalS = 60
	}

	var missing []string
	if configuration.Database.URL == "" {
		missing = append(missing, "DATABASE_URL")
	}
	if configuration.Redis.URL == "" {
		missing = append(missing, "REDIS_URL")
	}
	if configuration.Security.APIKey == "" {
		missing = append(missing, "API_KEY")
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("required configuration missing: %s", strings.Join(missing, ", "))
	}

	return configuration, nil
}

func lookupValue(lookup Lookup, key string) string {
	value, _ := lookup(key)
	return value
}

func lookupBool(lookup Lookup, key string) (bool, error) {
	value := lookupValue(lookup, key)
	if value == "" {
		return false, nil
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("%s must be a boolean: %w", key, err)
	}
	return parsed, nil
}

func lookupBoolWithDefault(lookup Lookup, key string, defaultValue bool) (bool, error) {
	value := lookupValue(lookup, key)
	if value == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("%s must be a boolean: %w", key, err)
	}
	return parsed, nil
}

func lookupInt(lookup Lookup, key string) (int, error) {
	value := lookupValue(lookup, key)
	if value == "" {
		return 0, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: %w", key, err)
	}
	return parsed, nil
}
