package fbauth

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func GetUIDByToken(idToken string) string {
	ctx, client := fbConnect()
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	return token.UID
}

func fbConnect() (context.Context, *auth.Client) {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing app: %v", err))
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Fb client: %n", err)
	}

	return ctx, client
}
