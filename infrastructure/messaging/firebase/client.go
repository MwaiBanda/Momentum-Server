package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/option"
)

const firebaseServiceTokenKey = "firebase"

func Open(
	ctx context.Context,
	tokenStore redis.Cmdable,
	projectID string,
) (*firebase.App, *messaging.Client, error) {
	token, err := tokenStore.Get(ctx, firebaseServiceTokenKey).Result()
	if err != nil {
		return nil, nil, fmt.Errorf("read Firebase service token from Redis: %w", err)
	}

	credentials := option.WithCredentialsJSON([]byte(token))
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID}, credentials)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize Firebase: %w", err)
	}

	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize Firebase messaging: %w", err)
	}

	log.Println("Connected to Firebase")
	return app, messagingClient, nil
}
