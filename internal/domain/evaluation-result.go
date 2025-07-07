package domain

type EvaluationResult struct {
	Score   int      `json:"score"`
	Level   string   `json:"level"`
	Summary []string `json:"summary"`
}
