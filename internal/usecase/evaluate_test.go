package usecase_test

import (
	"testing"

	"github.com/mateeusferro/fineval/internal/domain"
	"github.com/mateeusferro/fineval/internal/usecase"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		input    domain.EvaluationInput
		expected domain.EvaluationResult
	}{
		{
			name: "Ideal profile",
			input: domain.EvaluationInput{
				Income:     5000,
				Expenses:   2000,
				Debt:       1000,
				Savings:    10000,
				Dependents: 0,
			},
			expected: domain.EvaluationResult{
				Score: 100,
				Level: "High",
			},
		},
		{
			name: "High debt, low savings",
			input: domain.EvaluationInput{
				Income:     3000,
				Expenses:   2500,
				Debt:       2000,
				Savings:    1000,
				Dependents: 2,
			},
			expected: domain.EvaluationResult{
				Score: 76, // Calculation: 100 -20 +10 -10 -4 = 76
				Level: "Medium",
			},
		},
		{
			name: "Low score case",
			input: domain.EvaluationInput{
				Income:     2000,
				Expenses:   1900,
				Debt:       2000,
				Savings:    100,
				Dependents: 5,
			},
			expected: domain.EvaluationResult{
				Score: 70, // 100 -20 +10 -10 -10 = 70
				Level: "Medium",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := usecase.Evaluate(tt.input)

			if result.Score != tt.expected.Score {
				t.Errorf("Expected score %d, got %d", tt.expected.Score, result.Score)
			}
			if result.Level != tt.expected.Level {
				t.Errorf("Expected level %s, got %s", tt.expected.Level, result.Level)
			}
		})
	}
}
