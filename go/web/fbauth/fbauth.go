package fbauth

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func GetUIDByToken(idToken string) (string, error) {
	ctx, client, err := fbConnect()
	if err != nil {
		return "", err
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return "", entity.ErrInvalidIdToken
	}

	return token.UID, nil
}

func fbConnect() (context.Context, *auth.Client, error) {
	ctx := context.Background()

	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_SDK_CREDENTIALS")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
		return nil, nil, err
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
		return nil, nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error getting Fb client: %n", err)
		return nil, nil, err
	}

	return ctx, client, nil
}
