package invites

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strconv"
	"encoding/json"

	"../../../Projects/user-app-go/models/invites"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func parseInviteJSON() []invitemodel.Invite {
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
	var invites invitemodel.Invites

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'guests' which we defined above
	json.Unmarshal(byteValue, &invites)

	return invites.Invites
}

func parseDetailedInviteJSON() []invitemodel.DetailedInvite {
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
	var detailedInvites invitemodel.DetailedInvites

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'guests' which we defined above
	json.Unmarshal(byteValue, &detailedInvites)

	return detailedInvites.DetailedInvites
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", ReturnWatchlists)
	router.Get("/invite/{inviteID}", GetInvite)
	return router
}

func ReturnWatchlists(w http.ResponseWriter, r *http.Request) {
	invites := parseInviteJSON()
	render.JSON(w, r, invites)
}

func GetInvite(w http.ResponseWriter, r *http.Request) {
	inviteID, err := strconv.Atoi(chi.URLParam(r, "inviteID"))

	if err != nil {
		fmt.Println(err)
	}

	detailedInvites := parseDetailedInviteJSON()

	for i := range detailedInvites {
		if detailedInvites[i].ID == inviteID {
			render.JSON(w, r, detailedInvites[i])
		}
	}
}