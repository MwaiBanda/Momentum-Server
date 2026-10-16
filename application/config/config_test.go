package config

import (
	"strings"
	"testing"
)

func TestLoadConfigFromAppliesDefaultsAndOptionalValues(t *testing.T) {
	configuration, err := LoadConfigFrom(mapLookup(map[string]string{
		"DATABASE_URL": "postgres://database",
		"REDIS_URL":    "redis://cache",
		"API_KEY":      "secret",
	}))
	if err != nil {
		t.Fatalf("LoadConfigFrom() error = %v", err)
	}

	if configuration.Environment != "development" {
		t.Fatalf("Environment = %q, want %q", configuration.Environment, "development")
	}
	if configuration.Port != "8080" {
		t.Fatalf("Port = %q, want %q", configuration.Port, "8080")
	}
	if configuration.Firebase.ProjectID != "" {
		t.Fatalf("Firebase.ProjectID = %q, want empty", configuration.Firebase.ProjectID)
	}
	if configuration.SermonCloud.URL == "" {
		t.Fatal("SermonCloud.URL = empty, want default URL")
	}
}

func TestLoadConfigFromRequiresCoreValues(t *testing.T) {
	_, err := LoadConfigFrom(mapLookup(map[string]string{}))
	if err == nil {
		t.Fatal("LoadConfigFrom() error = nil, want missing configuration error")
	}

	for _, key := range []string{"DATABASE_URL", "REDIS_URL", "API_KEY"} {
		if !strings.Contains(err.Error(), key) {
			t.Errorf("LoadConfigFrom() error = %q, want %s", err, key)
		}
	}
}

func TestLoadConfigFromLoadsConfiguredValues(t *testing.T) {
	configuration, err := LoadConfigFrom(mapLookup(map[string]string{
		"ENVIRONMENT":               "production",
		"PORT":                      "9090",
		"DATABASE_URL":              "postgres://database",
		"REDIS_URL":                 "redis://cache",
		"API_KEY":                   "secret",
		"FIREBASE_PROJECT_ID":       "firebase-project",
		"SERMONCLOUD_URL":           "https://sermons.example.com/sermons",
		"STRIPE_SECRET_KEY":         "stripe-secret",
		"STRIPE_PUBLISHABLE_KEY":    "stripe-publishable",
		"STRIPE_VERSION":            "2024-01-01",
		"PLANNING_CENTER_URL":       "https://planning.example.com",
		"PLANNING_CENTER_TOKEN_URL": "https://planning.example.com/token",
		"EMAIL_APP_PASSWORD":        "email-password",
		"UNUSED_SETTING":            "ignored",
	}))
	if err != nil {
		t.Fatalf("LoadConfigFrom() error = %v", err)
	}

	if configuration.Stripe.PublishableKey != "stripe-publishable" {
		t.Fatalf("Stripe.PublishableKey = %q, want configured value", configuration.Stripe.PublishableKey)
	}
	if configuration.PlanningCenter.TokenURL != "https://planning.example.com/token" {
		t.Fatalf("PlanningCenter.TokenURL = %q, want configured value", configuration.PlanningCenter.TokenURL)
	}
	if configuration.SermonCloud.URL != "https://sermons.example.com/sermons" {
		t.Fatalf("SermonCloud.URL = %q, want configured value", configuration.SermonCloud.URL)
	}
	if configuration.Email.AppPassword != "email-password" {
		t.Fatalf("Email.AppPassword = %q, want configured value", configuration.Email.AppPassword)
	}
}

func mapLookup(values map[string]string) Lookup {
	return func(key string) (string, bool) {
		value, ok := values[key]
		return value, ok
	}
}
