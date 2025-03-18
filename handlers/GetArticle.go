package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetArticle(ctx *fiber.Ctx) error {
	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	articlesCollection := mongoDB.Collection("Articles")
	articleId := ctx.Params("id")

	var articleIdHex, err = primitive.ObjectIDFromHex(articleId)

	if err != nil {
		fmt.Println("issue parsing hex")
		return fiber.ErrInternalServerError
	}

	filter := bson.D{{"_id", articleIdHex}}

	var article types.Article

	err = articlesCollection.FindOne(context.Background(), filter).Decode(&article)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return fiber.ErrNotFound
	} else if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(article)
}
