package entity

type (
	User struct {
		UUID      string `json:"uuid"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		DOB       int64  `json:"dob"`
		ISDCode   string `json:"isd_code"`
		MobileNo  string `json:"mobile_no"`
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
	}
)
