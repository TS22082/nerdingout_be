package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	CreatedAt string             `json:"createdAt" bson:"createdAt"`
}

type BodyEntryType struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type Article struct {
	CreatorId   primitive.ObjectID `json:"creatorId" bson:"creatorId,omitempty"`
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Body        []BodyEntryType    `json:"body"  bson:"body"`
	CoverPhoto  string             `json:"coverPhoto,omitempty" bson:"coverPhoto,omitempty"`
	IsPublished bool               `json:"isPublished" bson:"isPublished"`
	CreatedAt   string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string             `json:"updatedAt" bson:"updatedAt"`
}

type HTTPRequestParams struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    interface{}
}

type T struct {
	Title       string `json:"title"`
	CoverPhoto  string `json:"coverPhoto"`
	Description string `json:"description"`
	Body        []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"body"`
	Published bool `json:"published"`
}
