package login

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", ReturnUser)
	router.Post("/", LoginUser)
	return router
}

func ReturnUser(w http.ResponseWriter, r *http.Request) {
	user := Login{
		Email: "test@test.com",
		Password: "Test123",
	}
	render.JSON(w, r, user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}

	response := make(map[string]string)
	response["message"] = "Logged in successfully"
	// response := map[string]string{
	// 	"Email": r.FormValue("Email"),
	// }
	render.JSON(w, r, response) // Return some demo response
}