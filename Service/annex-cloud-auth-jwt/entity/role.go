package entity

import "github.com/google/uuid"

type Role struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey;column:role_id"`
	Role string    `gorm:"type:varchar(15);not null;column:role_name"`
}
