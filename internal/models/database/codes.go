package databaseTables

import "gorm.io/gorm"

type Codes struct {
	gorm.Model
	ID         int    `json:"id" gorm:"primaryKey"`
	Code       string `json:"code" gorm:"unique"`
	Username   string `json:"username"`
	AssignRole string `json:"assign_role"`
	Used       bool   `json:"used"`
	UsedBy     string `json:"used_by,omitempty"`
}
