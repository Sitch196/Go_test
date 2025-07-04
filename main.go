package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Sitch196/Go_test/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file:", err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	collection = client.Database("golang_db").Collection("todos")

	handlers.SetCollection(collection)

	app := fiber.New()

	app.Get("/api/todos", handlers.GetTodos)
	// app.Post("/api/todos", handlers.CreateTodo)
	// app.Patch("/api/todos/:id", handlers.ToggleTodo)
	// app.Delete("/api/todos/:id", handlers.DeleteTodo)

	port := os.Getenv("PORT")

	fmt.Printf("Starting server on port %s\n", port)
	app.Listen(":" + port)
}
