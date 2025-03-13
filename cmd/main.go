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
	app := fiber.New()
	err := godotenv.Load()

	app.Use(middleware.MongoConnect(), middleware.Logging, middleware.CORS())

	articles := app.Group("/articles")
	articles.Get("/", middleware.VerifyToken, handlers.GetArticles)
	articles.Post("/", middleware.VerifyToken, handlers.PostArticle)
	articles.Patch("/:id", middleware.VerifyToken, handlers.PatchArticle)
	articles.Delete("/:id", middleware.VerifyToken, handlers.DeleteArticle)

	published := articles.Group("/published")
	published.Get("/", handlers.GetPublishedArticles)
	published.Get("/:id", handlers.GetPublishedArticle)

	auth := app.Group("/auth")
	auth.Get("gh", handlers.GhLogin)
	auth.Get("verify", middleware.VerifyToken, handlers.VerifyToken)

	const PORT = "0.0.0.0:8080"

	if err := app.Listen(PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	err = app.Shutdown()

	if err != nil {

		log.Fatal(err)
	}

	os.Exit(0)
}
