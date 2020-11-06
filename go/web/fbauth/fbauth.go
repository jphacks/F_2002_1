package fbauth

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func GetUIDByToken(idToken string) string {
	ctx, client := fbConnect()
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
	}

	return token.UID
}

func fbConnect() (context.Context, *auth.Client) {
	ctx := context.Background()

	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_SDK_CREDENTIALS")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error getting Fb client: %n", err)
	}

	return ctx, client
}
