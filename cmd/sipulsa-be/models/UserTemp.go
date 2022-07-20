package models

import (
	"database/sql"
	"time"
)

type UserTemp struct {
	tableName                struct{}       `pg:"user_temp"`
	Id                       int            `pg:"id"`
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
	Otp                      int            `pg:"otp"`
	OtpExpired               time.Time      `pg:"otp_expired"`
}
