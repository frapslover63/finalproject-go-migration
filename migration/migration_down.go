package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Down to remove collection from database
func Down(db *mongo.Database) {
	db.Collection("user").Drop(context.TODO())
	db.Collection("group").Drop(context.TODO())
	db.Collection("event").Drop(context.TODO())
	db.Collection("forum").Drop(context.TODO())
	db.Collection("game").Drop(context.TODO())
	db.Collection("objectUser").Drop(context.TODO())
}
