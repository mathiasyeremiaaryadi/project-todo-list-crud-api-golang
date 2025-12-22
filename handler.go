package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

	accessToken, err := GenerateAccessToken(loginUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	refreshToken, err := GenerateRefreshToken(loginUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshTokenHandler(c *fiber.Ctx) error {
	refreshTokenRequest := new(RefreshTokenRequest)
	err := c.BodyParser(refreshTokenRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	verifiedToken, err := VerifyToken(refreshTokenRequest.RefreshToken, os.Getenv("JWT_REFRESH_SECRET"), "refresh")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	jwtClaims := verifiedToken.Claims.(jwt.MapClaims)
	userId := int(jwtClaims["userId"].(float64))

	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": accessToken,
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
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	if existingTodo.UserID != authenticatedUserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
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
	todoId, _ := c.ParamsInt("id")

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)
	authenticatedUserId := int(authenticatedUser["userId"].(float64))

	err := DeleteTodo(todoId, authenticatedUserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func GetHandler(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	title := c.Query("title", "")
	isSorted, _ := strconv.ParseBool(c.Query("sort", "false"))

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)
	authenticatedUserId := int(authenticatedUser["userId"].(float64))
	todos, total, err := GetAllTodos(authenticatedUserId, page, limit, title, isSorted)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  todos,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}
