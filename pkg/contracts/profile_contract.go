package contracts

type (
	ProfileResponse struct {
		Email     string          `json:"email"`
		UserName  string          `json:"user_name"`
		FirstName string          `json:"first_name"`
		LastName  string          `json:"last_name"`
		Phone     PhoneResponse   `json:"phone"`
		Address   AddressResponse `json:"address"`
		UserType  string          `json:"user_type"`
	}

	WhoamiRequest struct {
		UserId string
	}
)
