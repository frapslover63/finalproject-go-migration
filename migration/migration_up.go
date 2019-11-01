package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// User struct for dummy data
type User struct {
	Email string
	Password string
	Username string
	ProfileImage string
	PhoneNumber string
	GameList []string
	EventList []string
}
type Comment struct {
	Username string
	ThreadComment string
}
type Thread struct {
	Name string
	Category string
	MakerUsername string
	// MakerImage string
	Description string
	CommentList []Comment
}
type Event struct {
	Name string
	Type string
	Games string
	Description string
	DateStart string
	DateEnd string
	Tag string
	Site string
	Location string
	Poster string
}

// Up to migrate database
func Up(db *mongo.Database) {
	// col := db.Collection("user")
	// col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	// col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
	mig1(db)
	mig2(db)
	mig3(db)
}

func mig1(db *mongo.Database){
	col := db.Collection("user")
	col.InsertOne(context.TODO(),
		User{
			Email: "test@gmail.com",
			Password: "password",
			Username: "TestUser",
			ProfileImage: "ProfileImage",
			PhoneNumber: "08123456789",
			GameList: []string{
				"Game1",
				"Game2",
				"Game3",
			},
			EventList: []string{
				"Event1",
				"Event2",
				"Event3",
			},
	})
}

func mig2(db *mongo.Database){
	col := db.Collection("thread")
	col.InsertOne(context.TODO(), 
	Thread{
			Name: "Thread 1",
			Category: "Category Thread 1",
			MakerUsername: "Maker Thread 1",
			Description: "Description Thread 1",
			CommentList: []Comment{
				{
					Username: "UserComment1",
					ThreadComment: "Comment1",
				},
				{
					Username: "UserComment2",
					ThreadComment: "Comment2",
				},
			},
	})
}

func mig3(db *mongo.Database){
	col := db.Collection("event")
	col.InsertOne(context.TODO(), 
		Event{
			Name: "Event 1",
			Type: "Type 1",
			Games: "Games 1",
			Description: "Description 1",
			DateStart: "DateStart 1",
			DateEnd: "DateEnd 1",
			Tag: "Tag 1",
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
		})
}
