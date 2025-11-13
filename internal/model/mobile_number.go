package model

type MobileNumber struct {
	// Number is the mobile number
	Number string `json:"number"`

	// Type indicates the type of mobile number (e.g., credit, debit)
	Type string `json:"type"`

	// IsActive indicates if the mobile number is active
	IsActive bool `json:"isActive"`

	// CountryCode is the country code associated with the mobile number
	CountryCode string `json:"countryCode"`
}

type AddMobileNumberRequest struct {
	MobileNumber MobileNumber `json:"mobileNumber"`
}

type AddMobileNumberFailureResponse struct {
	Error string `json:"error" example:"Failed to add mobile number"`
}
type AddMobileNumberSuccessResponse struct {
	MobileNumber MobileNumber `json:"mobileNumber"`
}

type DeleteMobileNumberSuccessResponse struct {
	Message  string `json:"message" example:"Mobile number deleted successfully"`
}