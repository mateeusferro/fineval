package config_test

import (
	"os"
	"testing"

	"github.com/mateeusferro/fineval/internal/config"
)

func TestLoadEnv(t *testing.T) {
	envContent := "TEST_KEY=test_value"
	err := os.WriteFile(".env", []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test .env file: %v", err)
	}
	defer os.Remove(".env")
	config.LoadEnv()

	val := config.EnvVariable("TEST_KEY")
	expected := "test_value"

	if val != expected {
		t.Errorf("Expected %s, got %s", expected, val)
	}
}
