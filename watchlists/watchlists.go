package watchlists

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strconv"
	"encoding/json"

	"../../../Projects/user-app-go/models/watchlists"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func parseJSON() []watchlistmodel.Watchlist {
	// Open our jsonFile
	jsonFile, err := os.Open("db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened db.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Guests array
	var watchlists watchlistmodel.Watchlists

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'guests' which we defined above
	json.Unmarshal(byteValue, &watchlists)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(watchlists.Watchlists); i++ {
		fmt.Println("Guest ID: " + strconv.Itoa(watchlists.Watchlists[i].ID))
		fmt.Println("Guest Name: " + watchlists.Watchlists[i].Name)
		fmt.Println("Guest Email: " + watchlists.Watchlists[i].Email)
		fmt.Println("Guest Notes: " + watchlists.Watchlists[i].Notes)
	}

	return watchlists.Watchlists
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", ReturnWatchlists)
	// router.Post("/", AddToWatchlist)
	return router
}

func ReturnWatchlists(w http.ResponseWriter, r *http.Request) {
	watchlists := parseJSON()
	render.JSON(w, r, watchlists)
}
