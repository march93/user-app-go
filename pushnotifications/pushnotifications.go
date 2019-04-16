package pushnotifications

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"google.golang.org/api/option"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	// router.Get("/", getEnv)
	router.Post("/", sendMessage)
	return router
}

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("guest-user-app-firebase-adminsdk-gi9qv-82f6209784.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

// func getEnv(w http.ResponseWriter, r *http.Request) {
// 	serverKey := os.Getenv("SERVER_KEY")
// 	render.JSON(w, r, serverKey)
// }

func sendMessage(w http.ResponseWriter, r *http.Request) {
	app := initializeAppWithServiceAccount()

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := os.Getenv("REGISTRATION_TOKEN")

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "WHAT'S COOKIN' HOMESLICE",
			Body:  "PHP COOKBOOK",
		},
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	render.JSON(w, r, response)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}
