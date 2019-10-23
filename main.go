package main

import (
	"context"
	"fmt"
	"os"

	"github.com/frapslover63/finalproject-go-migration/migration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Address address for mongodb host
	Address = "mongodb://localhost:27017"
	// DatabaseName name of database
	DatabaseName = "jwb"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing Argument (Up or Down)")
		return
	}
	arg := os.Args[1]
	clientOptions := options.Client().ApplyURI(Address)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := client.Database(DatabaseName)

	switch arg {
	case "up":
		fmt.Println("Migration Up")
		migration.Up(db)
	case "down":
		fmt.Println("Migration Down")
		migration.Down(db)
	}

	client.Disconnect(context.TODO())
}
