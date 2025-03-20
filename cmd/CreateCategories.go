package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	label string
}

func SeedCategoriesIfRequired() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/nerdingdout").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	categoryCollection := client.Database("nerdingout").Collection("Categories")

	var categories = make([]string, 6)
	categories[0] = "General"
	categories[1] = "Skateboarding"
	categories[2] = "Travel"
	categories[3] = "Programming"
	categories[4] = "Food / Beverage"
	categories[5] = "Technology"

	var count = 0
	for _, category := range categories {
		filter := bson.D{{"label", category}}
		var cat = new(Category)
		err = categoryCollection.FindOne(context.Background(), filter).Decode(&cat)

		if errors.Is(err, mongo.ErrNoDocuments) {
			count++

			_, err := categoryCollection.InsertOne(context.Background(), bson.M{
				"label": category,
			})

			if err != nil {
				fmt.Println("Cannot add")
			}
		}
	}

	fmt.Printf("Seeded %d categories\n", count)
	fmt.Printf("%d categories unchanged\n", len(categories)-count)
}
