package requests

import "database/sql"

type NewUserTemp struct {
	Name         string         `json:"name"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone_number"`
	ReferralCode sql.NullString `json:"referral_code"`
}
