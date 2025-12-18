package main

import "github.com/gofiber/fiber/v2"

func RegisterHandler(c *fiber.Ctx) error {
	registerRequest := new(UserRegisterRequest)
	err := c.BodyParser(registerRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"message": "invalid body request",
		})
	}

	hashedPassword, err := HashPassword(registerRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"message": "failed hash password",
		})
	}

	var user User
	user.Name = registerRequest.Name
	user.Email = registerRequest.Email
	user.Password = hashedPassword

	createdUser, err := CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"message": "failed register user",
		})
	}

	token, err := GenerateToken(createdUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"message": "faile tod generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]string{
		"token": token,
	})
}

func LoginHandler(c *fiber.Ctx) error {
	return nil
}

func CreateHandler(c *fiber.Ctx) error {
	return nil
}

func UpdateHandler(c *fiber.Ctx) error {
	return nil
}

func DeleteHandler(c *fiber.Ctx) error {
	return nil
}

func GetHandler(c *fiber.Ctx) error {
	return nil
}
