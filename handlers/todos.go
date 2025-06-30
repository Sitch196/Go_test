package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var todos []Todo

func GetTodos(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "Success",
		"data":   todos,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body can not be empty"})
	}
	todo.ID = len(todos) + 1
	todos = append(todos, *todo)

	return c.Status(201).JSON(fiber.Map{"status": "Success", "data": todo})
}

func ToggleTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i].Completed = !todo.Completed
			return c.Status(200).JSON(todos[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{"status": "failed", "message": "Todo not found"})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Todo deleted"})
		}
	}

	return c.Status(404).JSON(fiber.Map{"error ": "todo with this id not found"})
}
