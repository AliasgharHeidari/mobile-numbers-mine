package model

type User struct {
	// ID is the unique identifier for the user
	ID int64 `json:"id"`

	// Name is the first name of the user
	Name string `json:"name"`

	// FamilyName is the last name of the user
	FamilyName string `json:"familyName"`

	// Age is the age of the user
	Age int `json:"age"`

	// IsMarried indicates if the user is married
	IsMarried bool `json:"isMarried"`

	// MobileNumbers is a list of mobile numbers associated with the user
	MobileNumbers []MobileNumber `json:"mobileNumbers"`
}
