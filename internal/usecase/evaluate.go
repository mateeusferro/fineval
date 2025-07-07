package usecase

import (
	"github.com/mateeusferro/fineval/internal/domain"
)

func Evaluate(input domain.EvaluationInput) domain.EvaluationResult {
	score := 100
	summary := []string{}

	if (input.Debt / input.Income) > 0.5 {
		score -= 20
		summary = append(summary, "High debt to income ratio")
	}

	if (input.Income - input.Expenses) > 0 {
		score += 10
		summary = append(summary, "Positive monthly cash flow")
	}

	expensesThreeTimes := input.Expenses * 3
	if input.Savings >= expensesThreeTimes {
		score += 10
		summary = append(summary, "High savings buffer")
	} else {
		score -= 10
		summary = append(summary, "Low savings buffer")
	}

	score -= input.Dependents * 2

	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return domain.EvaluationResult{
		Score:   score,
		Level:   classifyScore(score),
		Summary: summary,
	}
}

func classifyScore(score int) string {
	switch {
	case score >= 80:
		return "High"
	case score >= 50:
		return "Medium"
	default:
		return "Low"
	}
}
