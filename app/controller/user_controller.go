package controller

import (
	"GoFiber-Login/app/models"
	"GoFiber-Login/config"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Render the register page
func RegisterPage(c *fiber.Ctx) error {
	// Serve the register page as a static file
	return c.SendFile("./public/register.html")
}

// Render the login page
func LoginPage(c *fiber.Ctx) error {
	// Serve the login page as a static file
	return c.SendFile("./public/login.html")
}

// Register a new user
func Register(c *fiber.Ctx) error {
	// Parse the user input from the request body
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Check if the user already exists in the database
	collection := config.MongoDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		// User already exists
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already registered. Please log in.",
		})
	}

	// Proceed only if the user is not already registered
	// Hash the password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user.Password = string(hashedPassword)

	// Save the user to the MongoDB collection
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Redirect to login page after successful registration
	return c.Redirect("/login")
}

// Login a user and issue a JWT token
func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse the login request input
	loginRequest := new(LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Find the user in the MongoDB collection
	collection := config.MongoDB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Create a JWT token for the logged-in user
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get JWT secret from the environment variable
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "JWT secret is not defined",
		})
	}
	fmt.Println("JWT Secret:", secret)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Return the signed JWT token in the response
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   signedToken,
	})
}
