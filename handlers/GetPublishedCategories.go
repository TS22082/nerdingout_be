package handlers

import (
	"errors"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPublishedCategories(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	categoriesCollection := mongoDB.Collection("Categories")
	articlesCollection := mongoDB.Collection("Articles")
	requestCtx := ctx.Context()

	var categories []types.Category
	cursor, err := categoriesCollection.Find(requestCtx, bson.M{})

	if err != nil {
		return fiber.ErrInternalServerError
	}
	if err := cursor.All(requestCtx, &categories); err != nil {
		return fiber.ErrInternalServerError
	}

	if len(categories) == 0 {
		return ctx.JSON([]types.Category{})
	}

	var response []types.Category
	for _, category := range categories {
		var article types.Article
		articleFilter := bson.M{"categoryId": category.Id, "isPublished": true}

		err = articlesCollection.FindOne(requestCtx, articleFilter).Decode(&article)
		if errors.Is(err, mongo.ErrNoDocuments) {
			continue
		} else if err != nil {
			return fiber.ErrInternalServerError
		}

		response = append(response, category)
	}

	return ctx.JSON(response)
}
