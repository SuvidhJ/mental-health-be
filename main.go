package main

import (
	"log"
	"os"

	"github.com/aditisaxena259/mental-health-be/config"
	"github.com/aditisaxena259/mental-health-be/models"
	"github.com/aditisaxena259/mental-health-be/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: No .env file found")
	}

	// Connect to PostgreSQL
	if err := config.ConnectDatabase(); err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}
	log.Println("✅ Connected to PostgreSQL!")

	models.AutoMigrateAll()
	models.SeedData()
	log.Println("📦 Database migrations completed successfully!")

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // frontend origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Serve local uploads (used when Cloudinary is not configured)
	app.Static("/uploads", "./uploads")

	// Setup API routes
	routes.SetupRoutes(app)

	// Dynamic port support
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local development
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
