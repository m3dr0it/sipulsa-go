package requests

type NewUserTemp struct {
	Name                     string `json:"name"`
	Username                 string `json:"username"`
	Email                    string `json:"email"`
	PhoneNumber              string `json:"phone_number"`
	RegisteredByReferralCode string `json:"registered_by_referral_code"`
}
