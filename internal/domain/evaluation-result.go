package domain

type EvaluationResult struct {
	Score   int      `json:"score"`
	Summary []string `json:"summary"`
}
