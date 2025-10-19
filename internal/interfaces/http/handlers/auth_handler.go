package handlers

import (
	"gohabits/internal/application/auth"
	"gohabits/internal/domain/user/dto"
	"gohabits/internal/infra"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService *auth.Service
}

func NewAuthHandler(container *infra.Container, jwtManager *auth.JWTManager) *AuthHandler {
	return &AuthHandler{
		userService: auth.NewAuthService(container, jwtManager),
	}
}

func (a *AuthHandler) Register(c *fiber.Ctx) error {

	var registerRequest dto.RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if errs := infra.ValidateStruct(registerRequest); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	createdUser, token, err := a.userService.RegisterUser(c, registerRequest)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message":      "User registered successfully",
		"user":         createdUser,
		"access_token": token,
	})
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {

	var loginRequest dto.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	res, err := a.userService.Login(c, loginRequest)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":      "Login successful",
		"access_token": res.AccessToken,
	})
}

func (a *AuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Convert float64 to uint (JWT claims store numbers as float64)
	userIDFloat, ok := userID.(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	authedUser, err := a.userService.Me(c, uint(userIDFloat))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user": authedUser,
	})
}
