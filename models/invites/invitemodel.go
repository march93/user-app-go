package invitemodel

import(
	"../../../../Projects/user-app-go/models/guests"
)

type Invites struct {
	Invites []Invite `json:"invites"`
}

type Invite struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Date int `json:"date"`
	Location string `json:"location"`
	Hosts []string `json:"hosts"`
	Watchlists []string `json:"watchlists"`
}

// Create a location model
// ID, Name
type Location struct {
	
}

type DetailedInvites struct {
	DetailedInvites []DetailedInvite `json:"detailedInvites"`
}

type DetailedInvite struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Company string `json:"company"`
	Country string `json:"country"`
	StartDate int `json:"start_date"`
	EndDate string `json:"end_date"`
	Location string `json:"location"`
	Hosts []guestmodel.Host `json:"hosts"`
	Watchlists []string `json:"watchlists"`
}