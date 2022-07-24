package requests

type NewUserTemp struct {
	Name                     string `json:"name" validate:"required"`
	Username                 string `json:"username" validate:"required"`
	Email                    string `json:"email" validate:"required,email"`
	PhoneNumber              string `json:"phone_number" validate:"required"`
	RegisteredByReferralCode string `json:"registered_by_referral_code"  `
}
