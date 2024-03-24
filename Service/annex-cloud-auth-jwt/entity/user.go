package entity

import (
	"regexp"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;column:user_id"`
	Username  string    `gorm:"type:varchar(50);not null;column:username"`
	Email     string    `gorm:"uniqueIndex;not null;column:email"`
	Password  string    `gorm:"not null;column:password"`
	RoleId    uuid.UUID `gorm:"column:role_id"`
	CreatedOn time.Time `gorm:"type:timestamptz;not null;column:created_on"`
	CreatedBy string    `gorm:"type:varchar(50);not null;column:created_by"`

	RoleRef Role `gorm:"foreignkey:RoleId;references:Id"`
}

func (u *User) ValidatePassword() bool {
	regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return regex.MatchString(u.Password)
}
