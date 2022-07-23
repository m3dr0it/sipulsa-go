package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserTemp struct {
	gorm.Model
	Id                       int            `gorm:PrimaryKey`
	Name                     string         `pg:"name"`
	Username                 string         `pg:"username"`
	PhoneNumber              string         `pg:"phone_number"`
	Email                    string         `pg:"email"`
	RegisteredByReferralCode sql.NullString `pg:"registered_by_referral_code"`
	CreatedAt                time.Time      `pg:"created_at"`
	CreatedBy                string         `pg:"created_by"`
	UpdatedAt                sql.NullTime   `pg:"updated_at"`
	UpdatedBy                sql.NullString `pg:"updated_by"`
	DeletedAt                sql.NullTime   `pg:"deleted_at"`
	DeletedBy                sql.NullString `pg:"deleted_by"`
	Otp                      string         `pg:"otp"`
	OtpExpired               time.Time      `pg:"otp_expired"`
}
