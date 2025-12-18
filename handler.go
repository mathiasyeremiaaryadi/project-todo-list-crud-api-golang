package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterHandler(c *fiber.Ctx) error {
	registerRequest := new(UserRegisterRequest)
	err := c.BodyParser(registerRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	hashedPassword, err := HashPassword(registerRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed hash password",
		})
	}

	var user User
	user.Name = registerRequest.Name
	user.Email = registerRequest.Email
	user.Password = hashedPassword

	err = CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
	})
}

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(UserLoginRequest)
	err := c.BodyParser(loginRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	loginUser, err := GetUser(loginRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid email or password",
		})
	}

	isLoginValid := VerifyPassword(loginUser.Password, loginRequest.Password)
	if !isLoginValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid email or password",
		})
	}

	token, err := GenerateToken(loginUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func CreateHandler(c *fiber.Ctx) error {
	todoRequest := new(TodoRequest)
	err := c.BodyParser(todoRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)

	var todo Todo
	todo.Title = todoRequest.Title
	todo.Description = todoRequest.Description
	todo.UserID = int(authenticatedUser["userId"].(float64))

	createedTodo, err := CreateTodo(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(NewTodoResponseFormat(createedTodo))
}

func UpdateHandler(c *fiber.Ctx) error {
	todoRequest := new(TodoRequest)
	err := c.BodyParser(todoRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)
	authenticatedUserId := int(authenticatedUser["userId"].(float64))

	todoId, _ := c.ParamsInt("id")
	existingTodo, err := GetTodo(todoId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	if existingTodo.UserID != authenticatedUserId {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	existingTodo.Title = todoRequest.Title
	existingTodo.Description = todoRequest.Description
	existingTodo.UserID = int(authenticatedUser["userId"].(float64))

	updatedTodo, err := UpdateTodo(existingTodo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(NewTodoResponseFormat(updatedTodo))
}

func DeleteHandler(c *fiber.Ctx) error {
	return nil
}

func GetHandler(c *fiber.Ctx) error {
	return nil
}
