package instructionals

import "go.mongodb.org/mongo-driver/bson/primitive"

type Instructional struct {
	ObjectId    primitive.ObjectID `json:"objectId" bson:"_id"`
	Id          string             `json:"id" bson:"id"`
	Title       string             `json:"title" bson:"title"`
	SearchTitle string             `json:"searchTitle" bson:"searchTitle"`
	Presenter   string             `json:"presenter" bson:"presenter"`
	Part        int                `json:"part" bson:"part"`
	Cover       string             `json:"cover" bson:"cover"`
	DownloadUrl string             `json:"download_url" bson:"download_url"`
	WatchUrl    string             `json:"watch_url" bson:"watch_url"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
}
