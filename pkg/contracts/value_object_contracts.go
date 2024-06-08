package contracts

type (
	AvailabilityResponse struct {
		Result bool `json:"result"`
	}

	DefaultResponse struct {
		Message string `json:"message"`
	}

	PhoneResponse struct {
		CountryCode string `json:"country_code"`
		Number      string `json:"number"`
	}

	ZipCodeResponse struct {
		Value string `json:"value"`
	}

	AddressResponse struct {
		Street  string          `json:"street"`
		Number  int             `json:"number"`
		ZipCode ZipCodeResponse `json:"zip_code"`
		City    string          `json:"city"`
	}
)
