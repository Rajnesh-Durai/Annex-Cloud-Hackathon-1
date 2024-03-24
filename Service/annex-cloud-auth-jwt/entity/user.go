package entity

import (
	"time"
    "regexp"
	"github.com/google/uuid"
)
type User struct {
	Id       uuid.UUID `db:"user_id" gorm:"type:uuid;primaryKey"`
	Username string    `db:"username" gorm:"type:varchar(255);not null"`
	Email    string `db:"email" gorm:"uniqueIndex;not null"`
	Password string `db:"password" gorm:"not null"`
	RoleId	 uuid.UUID `db:"role_id" gorm:"type:uuid;not null"`
	CreatedOn time.Time `db:"created_on" gorm:"type:timestamptz;not null"`
	CreatedBy string    `db:"created_by" gorm:"type:varchar(255);not null"`

	Role Role `gorm:"foreignKey:RoleId"`
}

func (u *User) ValidatePassword() bool {
    regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
    return regex.MatchString(u.Password)
}