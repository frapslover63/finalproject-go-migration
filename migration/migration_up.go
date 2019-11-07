package migration

import (
	"github.com/frapslover63/finalproject-go-migration/migration/data"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
)

// Up to migrate database
func Up(db *mongo.Database) {
	fmt.Println("User")
	data.UserData(db)
	fmt.Println("Thread")
	data.ThreadData(db)
	fmt.Println("Comment")
	data.CommentData(db)
	fmt.Println("Event")
	data.EventData(db)
	fmt.Println("Bookmark")
	data.BookmarkData(db)
}
