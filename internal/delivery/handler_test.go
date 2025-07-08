package delivery_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mateeusferro/fineval/internal/delivery"
	"github.com/mateeusferro/fineval/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestHandleEvaluate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	delivery.Routes(router)

	input := domain.EvaluationInput{
		Income:     5000,
		Expenses:   2000,
		Debt:       1000,
		Savings:    10000,
		Dependents: 0,
	}

	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/evaluate", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var result domain.EvaluationResult
	err := json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)

	assert.Equal(t, 100, result.Score)
	assert.Equal(t, "High", result.Level)
	assert.Contains(t, result.Summary, "Positive monthly cash flow")
}
