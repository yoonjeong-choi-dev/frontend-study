package docker_compose

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func ExecExample(address string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := Setup(ctx, address)
	if err != nil {
		return err
	}

	conn := db.Database("dockerCompose").Collection("example")
	vals := []interface{}{
		&User{Name: "yoonjeong", Age: 31},
		&User{Name: "yoonjeong-choi-dev", Age: 29},
	}

	if _, err := conn.InsertMany(ctx, vals); err != nil {
		return err
	}

	var user User
	if err := conn.FindOne(ctx, bson.M{"name": "yoonjeong"}).Decode(&user); err != nil {
		return err
	}

	if err := conn.Drop(ctx); err != nil {
		return err
	}

	fmt.Printf("Result Query: %#v\n", user)
	return nil
}
