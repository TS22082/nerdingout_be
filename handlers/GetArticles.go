package handlers

import (
	"context"
	"fmt"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetArticles get all articles published or not. Used for dashboard page on app
func GetArticles(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	articlesCollection := mongoDB.Collection("Articles")

	filter := bson.M{}

	var articles []types.Article

	cursor, err := articlesCollection.Find(context.Background(), filter)

	if err != nil {
		fmt.Println("Error finding articles")
		return fiber.ErrInternalServerError
	}

	if err := cursor.All(context.Background(), &articles); err != nil {
		fmt.Println("Error finding articles", err)
		return fiber.ErrInternalServerError
	}

	if len(articles) == 0 {
		return ctx.JSON([]map[string]interface{}{})
	}

	return ctx.JSON(articles)
}
