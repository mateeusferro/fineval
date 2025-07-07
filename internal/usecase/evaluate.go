package usecase

import "github.com/mateeusferro/fineval/internal/domain"

func Evaluate(input domain.EvaluationInput) domain.EvaluationResult {
	return domain.EvaluationResult{
		Score:   100,
		Summary: []string{"I don't know"},
	}
}
