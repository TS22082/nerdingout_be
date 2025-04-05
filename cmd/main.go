package main

import (
	"github.com/TS22082/nerdingout_be/handlers"
	"github.com/TS22082/nerdingout_be/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env: %v: ", err)
	}

	app := fiber.New()

	app.Use(middleware.MongoConnect(), middleware.Logging, middleware.CORS())

	articles := app.Group("/articles")
	published := articles.Group("/published")
	auth := app.Group("/auth")
	categories := app.Group("/categories")

	published.Get("/", handlers.GetPublishedArticles)
	published.Get("/:id", handlers.GetPublishedArticle)

	articles.Use(middleware.VerifyToken)

	articles.Get("/", handlers.GetArticles)
	articles.Get("/:id", handlers.GetArticle)
	articles.Post("/", handlers.PostArticle)
	articles.Patch("/:id", handlers.PatchArticle)
	articles.Delete("/:id", handlers.DeleteArticle)

	auth.Get("/gh", handlers.GhLogin)
	auth.Get("/verify", middleware.VerifyToken, handlers.VerifyToken)

	categories.Get("/", handlers.GetCategories)

	const PORT = "0.0.0.0:8080"

	if err := app.Listen(PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	if err := app.Shutdown(); err != nil {
		log.Fatal("Error shutting down server", err)
	}

	os.Exit(0)
}
