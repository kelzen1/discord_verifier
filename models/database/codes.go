package databaseTables

type Codes struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Username   string `json:"username"`
	AssignRole string `json:"assign_role"`
	Used       bool   `json:"used"`
	UsedBy     string `json:"used_by,omitempty"`
}
