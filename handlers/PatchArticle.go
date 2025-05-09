package handlers

import (
	"context"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// PatchArticle updates an article by its id, to be used in an admins dashboard
func PatchArticle(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	articleId := ctx.Params("id")

	if mongoDB == nil {
		return fiber.ErrInternalServerError
	}

	var articleIdHex, err = primitive.ObjectIDFromHex(articleId)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	articleCollection := mongoDB.Collection("Articles")
	filter := bson.D{{"_id", articleIdHex}}

	articleUpdates := types.Article{
		Title:       "",
		Body:        []types.BodyEntryType{},
		UpdatedAt:   time.Now().UTC().Format(time.RFC3339),
		IsPublished: false,
	}

	err = ctx.BodyParser(&articleUpdates)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	update := bson.M{"$set": articleUpdates}

	prevDoc := types.Article{
		CreatorId: primitive.ObjectID{},
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		Id:        primitive.ObjectID{},
	}

	err = articleCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&prevDoc)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	articleUpdates.CreatedAt = prevDoc.CreatedAt
	articleUpdates.CreatorId = prevDoc.CreatorId
	articleUpdates.Id = prevDoc.Id

	return ctx.Status(200).JSON(fiber.Map{
		"data": articleUpdates,
	})
}
