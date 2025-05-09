package handlers

import (
	"context"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// PostArticle creates a new unpublished article, to be used in an admins dashboard page
func PostArticle(ctx *fiber.Ctx) error {

	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	userId := ctx.Locals("userId").(string)
	timeNow := time.Now().UTC().Format(time.RFC3339)

	if mongoDB == nil {
		return ctx.Status(503).JSON(fiber.Map{
			"msg": "Database not available",
		})
	}

	userIdAsHex, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	var articleCollection = mongoDB.Collection("Articles")

	article := new(types.Article)

	article.UpdatedAt = timeNow
	article.CreatedAt = timeNow
	article.CreatorId = userIdAsHex

	err = ctx.BodyParser(&article)

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
