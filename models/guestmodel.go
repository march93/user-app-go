package guestmodel

// Guests struct which contains
// an array of guests
type Guests struct {
	Guests []Guest `json:"guests"`
}

// Guest struct
type Guest struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Date int `json:"date"`
	Location string `json:"location"`
	Company string `json:"company"`
	Hosts []Host `json:"hosts"`
	Level []string `json:"level"`
	ImageUrl string `json:"imageUrl"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	SignInData []SignInData `json:"signinData"`
	Notes string `json:"notes"`
}

type Host struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

type SignInData struct {
	ItemType string `json:"itemType"`
	Responses []Response `json:"responses"`
}

type Response struct {
	Question string `json:"question"`
	Response string `json:"response"`
}