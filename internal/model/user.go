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

type StatusUnauthorizedResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Authentication failed"`
}

type StatusNotFoundResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"User not found"`
}

type CreateUserSuccessResponse struct {
	// Message indicates the success message
	Message string `json:"message" example:"User created successfully"`

	// UserID is the ID of the created user
	UserID int `json:"user_id" example:"1"`
}

type CreateUserFailureResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Failed to create user"`
}

type LoginSuccessResponse struct {
	// Token is the JWT token
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type LoginFailureResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Authentication failed"`
}

type LoginRequest struct {
	// Username is the username of the user
	Username string `json:"username" example:"Aliasghar"`

	// Password is the password of the user
	Password string `json:"password" example:"1234"`
}

type GetUserListSuccessResponse struct {
	// Users is a list of users
	Users []User `json:"users"`
}

type GetUserListFailureResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Failed to retrieve user list"`
}

type GetUserByIDSuccessResponse struct {
	// User is the user object
	User User `json:"user"`
}

type UpdateUserByIDSuccessResponse struct {
	// Message indicates the success message
	Message string `json:"message" example:"User updated successfully"`
}
type UpdateUserByIDFailureResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Failed to update user"`
}

type UpdateUserByIDRequest struct {
	// Name is the first name of the user
	Name string `json:"name" example:"Amir"`

	// FamilyName is the last name of the user
	FamilyName string `json:"familyName" example:"Nigga"`

	// Age is the age of the user
	Age int `json:"age" example:"21"`

	// IsMarried indicates if the user is married
	IsMarried bool `json:"isMarried" example:"true"`
}

type DeleteUserByIDSuccessResponse struct {
	// Message indicates the success message
	Message string `json:"message" example:"User deleted successfully"`
}