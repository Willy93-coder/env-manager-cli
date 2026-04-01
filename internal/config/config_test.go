package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
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

	path := t.TempDir() + "/config.yaml"

	err := Save(cfg, path)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	cfgExpected, err := Load(path)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(cfgExpected, cfg) {
		t.Errorf("expected %v, got %v", cfg, cfgExpected)
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

func TestSave(t *testing.T) {
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

	path := t.TempDir() + "/config.yaml"

	err := Save(cfg, path)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		t.Fatalf("expected config file to be created at %s", path)
	}
	
	cfgExpected, err := Load(path)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(cfgExpected, cfg) {
		t.Errorf("expected %v, got %v", cfg, cfgExpected)
	}
}
