package main

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"mesage": err.Error()})
	}

	if err := CreateUser(db, &UserInput{SsoId: claims.User.Id, Name: claims.User.Name}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"mesage": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": token, "user": claims.User})
}
