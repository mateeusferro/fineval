package domain

type AuditlyInput struct {
	Actor    string                 `json:"actor"`
	Action   string                 `json:"action"`
	Resource string                 `json:"resource"`
	Metadata map[string]interface{} `json:"metadata"`
}
