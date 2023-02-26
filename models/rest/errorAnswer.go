package restModels

type ErrorAnswer struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
