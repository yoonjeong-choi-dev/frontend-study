package firebase

import (
	"context"
	fb "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

func Authenticate(ctx context.Context, collection string) (Client, error) {
	opt := option.WithCredentialsFile("tmp/service_account.json")
	app, err := fb.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init app")
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init firestore")
	}

	return &firebaseClient{Client: client, collection: collection}, nil
}
