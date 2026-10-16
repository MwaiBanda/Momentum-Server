package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const developmentEnvironment = "development"

type Config struct {
	Environment    string
	Port           string
	Database       DatabaseConfig
	Redis          RedisConfig
	Security       SecurityConfig
	Firebase       FirebaseConfig
	SermonCloud    SermonCloudConfig
	Stripe         StripeConfig
	PlanningCenter PlanningCenterConfig
	Email          EmailConfig
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

type Lookup func(string) (string, bool)

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && os.Getenv("ENVIRONMENT") == developmentEnvironment {
		return nil, fmt.Errorf("load .env file: %w", err)
	}

	return LoadConfigFrom(os.LookupEnv)
}

func LoadConfigFrom(lookup Lookup) (*Config, error) {
	configuration := &Config{
		Environment: lookupValue(lookup, "ENVIRONMENT"),
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
	}

	if configuration.Environment == "" {
		configuration.Environment = developmentEnvironment
	}
	if configuration.Port == "" {
		configuration.Port = "8080"
	}
	if configuration.SermonCloud.URL == "" {
		configuration.SermonCloud.URL = "https://api.sermoncloud.com/momentum-church-1/sermons"
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
