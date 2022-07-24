package models

import "time"

type User struct {
	ID                       int `gorm:"primaryKey"`
	MemberCode               string
	Username                 string
	Name                     string
	EncryptedPassword        string
	Address                  string
	PhoneNumber              string
	ReferralCode             string
	Email                    string
	RoleId                   int
	RegisteredByReferralCode string
	CreatedAt                time.Time
	CreatedBy                string
	UpdatedAt                time.Time
	UpdatedBy                string
	DeletedAt                time.Time
	DeletedBy                string
	IsActive                 bool
	IsDeleted                bool
}
