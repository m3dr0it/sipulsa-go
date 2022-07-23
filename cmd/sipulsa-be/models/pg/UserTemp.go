package models

import (
	"database/sql"
	"time"
)

type UserTemp struct {
	ID                       int `gorm:"primaryKey"`
	Name                     string
	Username                 string
	PhoneNumber              string
	Email                    string
	RegisteredByReferralCode string
	CreatedAt                time.Time
	CreatedBy                string
	UpdatedAt                sql.NullTime
	UpdatedBy                sql.NullString
	DeletedAt                sql.NullTime
	DeletedBy                sql.NullString
	Otp                      string
	OtpExpired               time.Time
}
