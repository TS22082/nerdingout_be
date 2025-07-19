// Package middleware provides various middleware functions for the dat board application,
// including MongoDB connection middleware to connect to MongoDB and make the client available to handlers.
package middleware

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnect is a middleware function that connects to a MongoDB instance.
// It retrieves the MongoDB URI from the environment variable MONGO_URI,
// establishes a connection to the MongoDB server, and pings the server to ensure connectivity.
// The MongoDB client is stored in the request context for subsequent handlers to use.
func MongoConnect() fiber.Handler {
	// Create a context with a timeout of 10 seconds for the MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve MongoDB URI from environment variables
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	mongo_uri := os.Getenv("MONGO_URI")

	fmt.Println("Mongo URI ==>", mongo_uri)

	clientOptions := options.Client().ApplyURI(mongo_uri).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		os.Exit(1)
	}

	// Ping the MongoDB server to ensure connectivity
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to ping MongoDB: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")

	// Middleware function to store the MongoDB client in the request context
	return func(c *fiber.Ctx) error {
		c.Locals("mongoDB", client.Database("nerdingout"))
		return c.Next()
	}
}
