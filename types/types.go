package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type Article struct {
	CreatorId   primitive.ObjectID `json:"creatorId" bson:"creatorId,omitempty"`
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Body        string             `json:"body" bson:"body"`
	IsPublished bool               `json:"isPublished" bson:"isPublished"`
	CreatedAt   string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string             `json:"updatedAt" bson:"updatedAt"`
}

type ArticleUpdate struct {
	Title     string `json:"title" bson:"title"`
	Body      string `json:"body" bson:"body"`
	UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
}
