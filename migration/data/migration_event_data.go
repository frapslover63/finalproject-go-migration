package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	DateStart int `bson:"dateStart"`
	DateEnd int `bson:"dateEnd"`
	Location string `bson:"location"`
	Poster string `bson:"poster"`
}

func EventData(db *mongo.Database) {
	col := db.Collection("event")
	col.InsertMany(context.TODO(), make([]interface{}, len([]Event{
		{
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
			DateStart: 1257894000,
			DateEnd: 1573436800,
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
		},
		{
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
			DateStart: 1257894000,
			DateEnd: 1573436800,
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
		},
	})))
}