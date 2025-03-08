package handlers

import (
	"context"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetArticles(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	articlesCollection := mongoDB.Collection("Articles")

	filter := bson.M{}
	var articles []types.Article

	cursor, err := articlesCollection.Find(context.Background(), filter)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	if err := cursor.All(context.Background(), &articles); err != nil {
		return fiber.ErrInternalServerError
	}

	response := fiber.Map{
		"msg": articles,
	}

	return ctx.JSON(response)
}
