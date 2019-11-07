package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)
// struct for dummy data
type User struct {
	Email string `bson:"email"`
	Password string `bson:"password"`
	Username string `bson:"username"`
	ProfileImage string `bson:"profileImage"`
	PhoneNumber string `bson:"phoneNumber"`
	GameList []string `bson:"gameList"`
}

func UserData(db *mongo.Database){
	col := db.Collection("user")
	col.InsertOne(context.TODO(),
		User{
			Email: "test@gmail.com",
			Password: "password",
			Username: "TestUser",
			ProfileImage: "ProfileImage",
			PhoneNumber: "08123456789",
			GameList: []string{
				"Dota",
				"PUBG",
				"YuGiOh",
			},
		})
	col.InsertOne(context.TODO(),

		User{
			Email: "admin@gmail.com",
			Password: "abcdef",
			Username: "Admin",
			ProfileImage: "ProfileImageAdmin",
			PhoneNumber: "08123456788",
			GameList: []string{
				"Dota",
			},
		})
}
