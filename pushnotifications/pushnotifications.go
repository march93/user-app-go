package pushnotifications

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
	router.Post("/data", sendDataMessage)
	router.Post("/notify", sendMessage)
	router.Post("/subscribeToTopic", subscribeToTopic)
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

func sendDataMessage(w http.ResponseWriter, r *http.Request) {
	app := initializeAppWithServiceAccount()
	topic := "Notifications"

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title": "Por Favor",
			"body":  "Venha na minhas festa",
		},
		Topic: topic,
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

func sendMessage(w http.ResponseWriter, r *http.Request) {
	app := initializeAppWithServiceAccount()
	topic := "Notifications"

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		Notification: &messaging.Notification{
			Title: "WHAT'S COOKIN' HOMESLICE",
			Body:  "PHP COOKBOOK",
		},
		Data: map[string]string{
			"title": "Hola",
			"body":  "Senõr",
		},
		Topic: topic,
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

func subscribeToTopic(w http.ResponseWriter, r *http.Request) {
	app := initializeAppWithServiceAccount()

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	topic := "Notifications"

	r.ParseForm()
	registrationTokens := []string{
		r.FormValue("token"),
	}
	response, err := client.SubscribeToTopic(ctx, registrationTokens, topic)
	if err != nil {
		log.Fatalln(err)
	}

	// See the TopicManagementResponse reference documentation
	// for the contents of response.
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")

	// response := make(map[string]bool)
	// response["success"] = true

	render.JSON(w, r, map[string]bool{"success": true})
}
