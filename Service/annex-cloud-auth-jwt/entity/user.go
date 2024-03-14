package entity

import (
	"time"

	"github.com/google/uuid"
)
type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"type:nvarchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role 	 string `gorm:"type:nvarchar(50);not null"`
	CreatedOn time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy string    `gorm:"type:nvarchar(255);not null"`
}