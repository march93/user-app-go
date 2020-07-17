package pushnotifications

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"google.golang.org/api/option"
)

// Routes calls related to push notifications
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/data", sendDataMessage)
	router.Post("/subscribeToTopic", subscribeToTopic)
	router.Post("/unsubscribeFromTopic", unsubscribeFromTopic)
	return router
}

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("guest-mobile-dev-bda13-firebase-adminsdk-53xsi-eaa09a3e4e.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

func sendDataMessage(w http.ResponseWriter, r *http.Request) {
	app := initializeAppWithServiceAccount()

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// Payload constants
	topic := "Notifications"
	title := "Psst"
	body := "Leave your laptop unlocked just for today"
	imgURL := "https://i.imgur.com/4ABlT4r.jpg"
	action := "Acknowledge"
	guestID := "9"
	guestName := "Shantell Crooks"
	screen := "guestDetail"

	message := &messaging.Message{
		Data: map[string]string{
			"title":     title,
			"body":      body,
			"imgURL":    imgURL,
			"action":    action,
			"guestID":   guestID,
			"guestName": guestName,
			"screen":    screen,
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					ContentAvailable: true,
					Alert: &messaging.ApsAlert{
						Title: title,
						Body:  body,
					},
					Category: "NOTIFICATION_CATEGORY",
					CustomData: map[string]interface{}{
						"imgURL":    imgURL,
						"action":    action,
						"guestID":   guestID,
						"guestName": guestName,
						"screen":    screen,
					},
					MutableContent: true,
				},
			},
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
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]bool{"success": false})
		return
	}

	// See the TopicManagementResponse reference documentation
	// for the contents of response.
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]bool{"success": true})
}

func unsubscribeFromTopic(w http.ResponseWriter, r *http.Request) {
	app := initializeAppWithServiceAccount()

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	topic := "Notifications"

	r.ParseForm()
	registrationTokens := []string{
		r.FormValue("token"),
	}
	response, err := client.UnsubscribeFromTopic(ctx, registrationTokens, topic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]bool{"success": false})
		return
	}

	// See the TopicManagementResponse reference documentation
	// for the contents of response.
	fmt.Println(response.SuccessCount, "tokens were unsubscribed successfully")
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]bool{"success": true})
}
