package handlers

import (
	"context"
	"errors"

	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetArticle gets an article by an article id for a single article edit/ view admin page
func GetArticle(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	articlesCollection := mongoDB.Collection("Articles")
	articleId := ctx.Params("id")

	var articleIdHex, err = primitive.ObjectIDFromHex(articleId)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	filter := bson.D{{Key: "_id", Value: articleIdHex}}

	var article types.Article

	err = articlesCollection.FindOne(context.Background(), filter).Decode(&article)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return fiber.ErrNotFound
	} else if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(article)
}
