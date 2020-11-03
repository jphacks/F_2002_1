package infrastracture

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/jphacks/F_2002_1/go/interface/myfirebase"
)

type FbHandler struct {
	Ctx    context.Context
	Client *auth.Client
}

func NewFbHandler() myfirebase.FbHandler {
	fbHandler := new(FbHandler)
	fbHandler.Ctx, fbHandler.Client = fbConnect()
	return fbHandler
}

func (handler *FbHandler) GetUIDByToken(idToken string) string {
	token, err := handler.Client.VerifyIDToken(handler.Ctx, idToken)
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
