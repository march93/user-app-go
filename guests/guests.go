package guests

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"

	"../../../Projects/user-app-go/models/guests"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func parseJSON() []guestmodel.Guest {
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
	var guests guestmodel.Guests

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'guests' which we defined above
	json.Unmarshal(byteValue, &guests)

	return guests.Guests
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", ReturnGuests)
	// router.Post("/", AddGuest)
	return router
}

func ReturnGuests(w http.ResponseWriter, r *http.Request) {
	guests := parseJSON()
	render.JSON(w, r, guests)
}
