package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)
// struct for dummy data
type Thread struct {
	Name string `bson:"name"`
	Category string `bson:"category"`
	MakerUsername string `bson:"makerUsername"`
	Description string `bson:"description"`
}

func ThreadData(db *mongo.Database) {
	col := db.Collection("thread")
	col.InsertMany(context.TODO(), make([]interface{}, len([]Thread{
		{
			Name:          "Thread 1",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 1",
		},
		{
			Name:          "Thread 2",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 2",
		},
		{
			Name:          "Thread 3",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 3",
		},
		{
			Name:          "Thread 4",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 4",
		},
		{
			Name:          "Thread 5",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 5",
		},
		{
			Name:          "Thread 6",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 6",
		},
		{
			Name:          "Thread 7",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 7",
		},
		{
			Name:          "Thread 8",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 8",
		},
		{
			Name:          "Thread 9",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 9",
		},
		{
			Name:          "Thread 10",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 10",
		},
		{
			Name:          "Thread 11",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 11",
		},
		{
			Name:          "Thread 12",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 12",
		},
		{
			Name:          "Thread 13",
			Category:      "Shooting",
			MakerUsername: "TestUser",
			Description:   "Description Thread 13",
		},
		{
			Name:          "Thread 14",
			Category:      "TCG",
			MakerUsername: "Admin",
			Description:   "Description Thread 14",
		},
	})))
}