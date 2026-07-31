package model

import "gorm.io/gorm"

// Topic represents the syllabus learning topics stored in PostgreSQL.
type Topic struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
}
