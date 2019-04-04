package watchlistmodel

type Watchlists struct {
	Watchlists []Watchlist `json:"watchlists"`
}

type Watchlist struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Notes string `json:"notes"`
	Level []string `json:"level"`
	Aliases []string `json:"aliases"`
	ImageUrl string `json:"imageUrl"`	
}