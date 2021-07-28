package dto

type (
	UserOnboardingRequest struct {
		DeviceID   string
		DeviceType string
	}

	UpdateProfileRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		DOB       int64  `json:"dob"`
	}
)
