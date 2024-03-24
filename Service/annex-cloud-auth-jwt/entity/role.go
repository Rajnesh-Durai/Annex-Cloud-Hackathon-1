package entity

import "github.com/google/uuid"

type Role struct {
	Id   uuid.UUID `db:"role_id" gorm:"type:uuid;primaryKey"`
	Role string    `db:"role_name" gorm:"type:varchar(255);not null"`
}
