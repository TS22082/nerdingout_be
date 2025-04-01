package main

import (
	"github.com/TS22082/nerdingout_be/handlers"
	"github.com/TS22082/nerdingout_be/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Use(middleware.MongoConnect(), middleware.Logging, middleware.CORS())

	articles := app.Group("/articles")
	published := articles.Group("/published")
	auth := app.Group("/auth")
	categories := app.Group("/categories")

	published.Get("/", handlers.GetPublishedArticles)
	published.Get("/:id", handlers.GetPublishedArticle)

	articles.Get("/", middleware.VerifyToken, handlers.GetArticles)
	articles.Get("/:id", middleware.VerifyToken, handlers.GetArticle)
	articles.Post("/", middleware.VerifyToken, handlers.PostArticle)
	articles.Patch("/:id", middleware.VerifyToken, handlers.PatchArticle)
	articles.Delete("/:id", middleware.VerifyToken, handlers.DeleteArticle)

	auth.Get("/gh", handlers.GhLogin)
	auth.Get("/verify", middleware.VerifyToken, handlers.VerifyToken)

	categories.Get("/", handlers.GetCategories)

	const PORT = "0.0.0.0:8080"

	if err := app.Listen(PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
