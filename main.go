package main

import (
	"log"
	"os"

	"github.com/Sitch196/Go_test/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env")
	}

	PORT := os.Getenv("PORT")

	// Public route
	app.Get("/", handlers.GetTodos)

	// Auth route
	app.Post("/login", handlers.Login)

	// Protected routes
	app.Post("/api/todos", handlers.JWTProtected(), handlers.CreateTodo)
	app.Patch("/api/todos/:id", handlers.JWTProtected(), handlers.ToggleTodo)
	app.Delete("/api/todos/:id", handlers.JWTProtected(), handlers.DeleteTodo)

	log.Fatal(app.Listen(":" + PORT))
}
