package handlers

import (
	"context"
	"fmt"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPublishedArticles get all the published articles for the main page
func GetPublishedArticles(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	articlesCollection := mongoDB.Collection("Articles")
	categoryId := ctx.Query("categoryId")

	fmt.Println(categoryId)

	filter := bson.D{{"isPublished", true}}

	if categoryId != "" {
		categoryAsHex, err := primitive.ObjectIDFromHex(categoryId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		filter = append(filter, bson.E{Key: "categoryId", Value: categoryAsHex})
	}

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
