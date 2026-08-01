package model

import "time"

// Topic represents the syllabus learning topics stored in PostgreSQL.
type Topic struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
