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
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 1",
			Category:      "Shooting",
			MakerUsername: "TestUser",
			Description:   "Description Thread 1",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 2",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 2",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 3",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 3",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 4",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 4",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 5",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 5",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 6",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 6",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 7",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 7",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 8",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 8",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 9",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 9",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 10",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 10",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 11",
			Category:      "Moba",
			MakerUsername: "Admin",
			Description:   "Description Thread 11",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 12",
			Category:      "Moba",
			MakerUsername: "TestUser",
			Description:   "Description Thread 12",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 13",
			Category:      "Shooting",
			MakerUsername: "TestUser",
			Description:   "Description Thread 13",
		})
	col.InsertOne(context.TODO(),
		Thread{
			Name:          "Thread 14",
			Category:      "TCG",
			MakerUsername: "Admin",
			Description:   "Description Thread 14",
		})
}