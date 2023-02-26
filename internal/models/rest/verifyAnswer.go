package restModels

type VerifyAnswer struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
}
