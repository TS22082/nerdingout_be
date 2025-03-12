package handlers

import (
	"context"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func PostArticle(ctx *fiber.Ctx) error {

	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	timeNow := time.Now().UTC().Format(time.RFC3339)

	if mongoDB == nil {
		return ctx.Status(503).JSON(fiber.Map{
			"msg": "Database not available",
		})
	}

	var articleCollection = mongoDB.Collection("Articles")

	article := new(types.Article)

	article.UpdatedAt = timeNow
	article.CreatedAt = timeNow

	err := ctx.BodyParser(&article)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"msg": "Issue parsing json",
		})
	}

	res, err := articleCollection.InsertOne(context.Background(), article)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"msg": "Failed to insert article",
		})
	}

	article.Id = res.InsertedID.(primitive.ObjectID)

	response := fiber.Map{
		"data": article,
	}

	return ctx.JSON(response)
}
