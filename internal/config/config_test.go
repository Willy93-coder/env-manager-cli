package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if cfg.Database.Host == "" {
		t.Error("expected database host to be set")
	}
}

func TestDSN(t *testing.T) {
	cfg := &Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "test",
			Password: "pass",
			DBName:   "testdb",
			SSLMode:  "disable",
		},
	}

	dsn := cfg.DSN()
	expected := "host=localhost port=5432 user=test password=pass dbname=testdb sslmode=disable"
	if dsn != expected {
		t.Errorf("expected %q, got %q", expected, dsn)
	}
}
