package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteArticle(ctx *fiber.Ctx) error {
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
	filter := bson.M{"_id": articleIdHex}

	_, err = articleCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := fiber.Map{
		"msg": "Article was deleted",
	}

	return ctx.JSON(response)
}
