package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)
// struct for dummy data
type Comment struct {
	MasterThreadId string  `bson:"masterThreadId"`
	Username string `bson:"username"`
	ThreadComment string `bson:"threadComment"`
}

func CommentData(db *mongo.Database) {
	col := db.Collection("comment")
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "TestUser",
			ThreadComment: "Comment 1.1",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 1.2",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "TestUser",
			ThreadComment: "Comment 1.3",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 2.1",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 2.2",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "TestUser",
			ThreadComment: "Comment 1.4",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 1.5",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 1.6",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "TestUser",
			ThreadComment: "Comment 1.7",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 1.8",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "TestUser",
			ThreadComment: "Comment 1.9",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 1.10",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "Admin",
			ThreadComment: "Comment 1.11",
		})
	col.InsertOne(context.TODO(),
		Comment{
			MasterThreadId: "",
			Username: "TestUser",
			ThreadComment: "Comment 2.3",
		})
}