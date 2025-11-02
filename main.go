package main

import (
	"adhayan-go-crm-basic/database"
	"adhayan-go-crm-basic/lead"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead) // <-- add :id here
}

func initDBConnection() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL not set in .env or environment")
	}

	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	database.DBConn.AutoMigrate(&lead.Lead{})

	log.Println("✅ Database connected successfully!")
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	initDBConnection()
	app.Listen(":3000")
	sqlDB, _ := database.DBConn.DB()
	defer sqlDB.Close()
}
