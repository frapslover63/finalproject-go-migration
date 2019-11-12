package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)
// struct for dummy data
type Event struct {
	Name string `bson:"name"`
	MakerUsername string `bson:makerUsername`
	Type string `bson:"type"`
	Games []string `bson:"games"`
	Category []string `bson:"category"`
	Description string `bson:"description"`
	Site string `bson:"site"`
	DateStart time.Time `bson:"dateStart"`
	DateEnd time.Time `bson:"dateEnd"`
	Location string `bson:"location"`
	Poster string `bson:"poster"`
}

func EventData(db *mongo.Database) {
	col := db.Collection("event")
	col.InsertOne(context.TODO(),
		Event{
			Name: "Event 1",
			MakerUsername: "TestUser",
			Type: "Competition",
			Games: []string{
				"Dota",
				"PUBG",
			},
			Category: []string{
				"Moba",
				"Shooting",
			},
			Description: "Description 1",
			DateStart: time.Date(2019, time.November, 10, 12, 0, 1, 0, time.UTC),
			DateEnd: time.Date(2019, time.November, 12, 12, 0, 1, 0, time.UTC),
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
		})
	col.InsertOne(context.TODO(),
		Event{
			Name: "Event 2",
			MakerUsername: "Admin",
			Type: "Event",
			Games: []string{
				"YuGiOh",
			},
			Category: []string{
				"TCG",
			},
			Description: "Description 1",
			DateStart: time.Date(2019, time.November, 15, 12, 0, 1, 0, time.UTC),
			DateEnd: time.Date(2019, time.November, 20, 12, 0, 1, 0, time.UTC),
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
		})
}