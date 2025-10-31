package model

type User struct {
	// ID is the unique identifier for the user
	ID int `json:"id"`

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

type CreateUserRequest struct {
	// Name is the first name of the user
	Name string `json:"name" example:"John"`

	// FamilyName is the last name of the user
	FamilyName string `json:"familyName" example:"Doe"`

	// Age is the age of the user
	Age int `json:"age" example:"30"`

	// IsMarried indicates if the user is married
	IsMarried bool `json:"isMarried" example:"false"`
}

type CreateUserSuccessResponse struct {
	// Message indicates the success message
	Message string `json:"message" example:"User created successfully"`

	// UserID is the ID of the created user
	UserID int `json:"user_id" example:"1"`
}
