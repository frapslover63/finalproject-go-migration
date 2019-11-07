package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Down to remove collection from database
func Down(db *mongo.Database) {
	db.Collection("user").Drop(context.TODO())
	db.Collection("thread").Drop(context.TODO())
	db.Collection("event").Drop(context.TODO())
	db.Collection("comment").Drop(context.TODO())
	db.Collection("bookmark").Drop(context.TODO())
}
