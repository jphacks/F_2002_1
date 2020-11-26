package myfirebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func initFirebase(conf *firebase.Config) (context.Context, *firebase.App, error) {
	ctx := context.Background()

	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_SDK_CREDENTIALS")))
	if err != nil {
		log.Printf("error credentials from json: %v\n", err)
		return nil, nil, err
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
		return nil, nil, err
	}

	return ctx, app, nil
}

func initAuth() (context.Context, *auth.Client, error) {
	conf := &firebase.Config{
		DatabaseURL: "https://speak-vegetable.firebaseio.com/",
	}
	ctx, app, err := initFirebase(conf)
	if err != nil {
		return nil, nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error getting Fb client: %n", err)
		return nil, nil, err
	}

	return ctx, client, nil
}

func initDatabase() (context.Context, *db.Client, error) {
	ctx, app, err := initFirebase(nil)
	if err != nil {
		return nil, nil, err
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return ctx, client, nil
}

func GetUIDByToken(idToken string) (string, error) {
	ctx, client, err := initAuth()
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
