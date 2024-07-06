package models

import "gorm.io/gorm"

type Column struct {
	gorm.Model        // Встроенная модель GORM, которая добавляет поля ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `gorm:"unique;not null"`
	Order      uint   `gorm:"primaryKey;autoIncrement"`
}
