package models

import "github.com/google/uuid"

type Users struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"-"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email" `
	Password  string    `json:"password" gorm:"not null" validate:"required,min=6,max=15"`
	FirstName string    `json:"firstname" gorm:"not null" validate:"required"`
	LastName  string    `json:"lastname" gorm:"not null" validate:"required"`
	Role      string    `json:"-" gorm:"not null"`
	CreatedAt int       `json:"created_at" gorm:"not null"`
	UpdatedAt int       `json:"updated_at" gorm:"not null"`
}
