package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)
// struct for dummy data
type Comment struct {
	MasterThreadId int  `bson:"masterThreadId"`
	Username string `bson:"username"`
	ThreadComment string `bson:"threadComment"`
}

func CommentData(db *mongo.Database) {
	col := db.Collection("comment")
	col.InsertMany(context.TODO(), make([]interface{}, len([]Comment{
		{
			MasterThreadId: 0,
			Username: "TestUser",
			ThreadComment: "Comment 1.1",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 1.2",
		},
		{
			MasterThreadId: 0,
			Username: "TestUser",
			ThreadComment: "Comment 1.3",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 2.1",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 2.2",
		},
		{
			MasterThreadId: 0,
			Username: "TestUser",
			ThreadComment: "Comment 1.4",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 1.5",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 1.6",
		},
		{
			MasterThreadId: 0,
			Username: "TestUser",
			ThreadComment: "Comment 1.7",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 1.8",
		},
		{
			MasterThreadId: 0,
			Username: "TestUser",
			ThreadComment: "Comment 1.9",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 1.10",
		},
		{
			MasterThreadId: 0,
			Username: "Admin",
			ThreadComment: "Comment 1.11",
		},
		{
			MasterThreadId: 0,
			Username: "TestUser",
			ThreadComment: "Comment 2.3",
		},
	})))
}