package entity

import "time"

type User struct {
	Id                  string    `gorm:"primaryKey;column:id;"`
	IsActive            int       `gorm:"column:is_active;"`
	VerificationDueDate time.Time `gorm:"column:verification_due_date;"`
	NotVerification     int       `gorm:"column:not_verification;"`
	NotVerificationDate time.Time `gorm:"column:not_verification;"`
}

func (User) TableName() string {
	return "users"
}
