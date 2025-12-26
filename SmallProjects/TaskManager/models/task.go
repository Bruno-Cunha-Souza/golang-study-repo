package models

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string `gorm:"default:pendente"`
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
