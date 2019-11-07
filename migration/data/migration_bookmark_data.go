package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)
// struct for dummy data
type Bookmark struct {
	UserId int `bson:"userId"`
	EventId int `bson:"eventId"`
}

func BookmarkData(db *mongo.Database) {
	col := db.Collection("bookmark")
	col.InsertMany(context.TODO(), make([]interface{}, len([]Bookmark{
		{
			UserId: 0,
			EventId: 0,
		},
	})))
}