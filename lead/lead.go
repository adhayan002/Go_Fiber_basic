package lead

import (
	"adhayan-go-crm-basic/database"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   string
}

// Get all leads
func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

// Get single lead by ID
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead

	result := db.First(&lead, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Lead not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Database error")
	}

	return c.JSON(lead)
}

// Create new lead
func CreateLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Create(&lead)
	return c.JSON(lead)
}

// Delete lead by ID
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead

	result := db.First(&lead, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Lead not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Database error")
	}

	db.Delete(&lead)
	return c.SendString("Lead deleted successfully")
}
