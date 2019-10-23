package migration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// User struct for dummy data
type User struct {
	Email string
}

// Up to migrate database
func Up(db *mongo.Database) {
	col := db.Collection("user")
	col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
	col = db.Collection("group")
	col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
	col = db.Collection("event")
	col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
	col = db.Collection("forum")
	col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
	col = db.Collection("game")
	col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
	col = db.Collection("objectUser")
	col.InsertOne(context.TODO(), User{Email: "test@gmail.com"})
	col.DeleteOne(context.TODO(), User{Email: "test@gmail.com"})
}
