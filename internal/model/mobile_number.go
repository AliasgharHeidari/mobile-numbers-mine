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
