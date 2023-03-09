package databaseTables

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
	Role string `json:"role"`
}
