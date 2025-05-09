package handlers

import (
	"context"
	"fmt"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCategories retrieves all the categories available in the DB
func GetCategories(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	categoriesCollection := mongoDB.Collection("Categories")

	filter := bson.M{}

	var categories []types.Category

	cursor, err := categoriesCollection.Find(context.Background(), filter)

	if err != nil {
		fmt.Println("Error finding categories", err)
		return fiber.ErrInternalServerError
	}

	if err := cursor.All(context.Background(), &categories); err != nil {

		fmt.Println("Error getting categories", err)
	}

	if len(categories) == 0 {
		return ctx.JSON([]map[string]interface{}{})
	}

	return ctx.JSON(categories)
}
