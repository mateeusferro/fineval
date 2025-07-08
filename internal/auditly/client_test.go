package auditly_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mateeusferro/fineval/internal/auditly"
	"github.com/mateeusferro/fineval/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env: %v", err)
	}
	mockedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/log", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var receivedData domain.AuditlyInput
		err := json.NewDecoder(r.Body).Decode(&receivedData)
		assert.NoError(t, err)
		assert.NotNil(t, receivedData)

		w.WriteHeader(http.StatusCreated)
	}))

	originalHost := os.Getenv("AUDITLY_HOST")
	defer func() {
		os.Setenv("AUDITLY_HOST", originalHost)
	}()
	os.Setenv("AUDITLY_HOST", mockedServer.URL)

	logData := domain.AuditlyInput{
		Actor:    "mockActor",
		Action:   "MOCK_ACTION",
		Resource: "mockResource",
		Metadata: map[string]interface{}{
			"mock": "Mock",
		},
	}

	success := auditly.Send(logData)

	assert.True(t, success)

	mockedServer.Close()
}
