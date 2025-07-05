package domain

type EvaluationInput struct {
	Income     float64 `json:"income"`
	Expenses   float64 `json:"expenses"`
	Debt       float64 `json:"debt"`
	Savings    float64 `json:"savings"`
	Dependents int     `json:"dependents"`
}
