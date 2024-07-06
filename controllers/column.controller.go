package controllers

import (
	"awesomeProject5/initializers"
	"awesomeProject5/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CreateColumnHandler(c *fiber.Ctx) error {
	var payload *models.Column

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newColumn := models.Column{
		Title: payload.Title,
		Order: payload.Order,
	}

	result := initializers.DB.Create(&newColumn)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"column": newColumn}})
}

func FindColumns(c *fiber.Ctx) error {
	var columns []models.Column
	results := initializers.DB.Find(&columns)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "columns": columns})
}

func DeleteColumn(c *fiber.Ctx) error {
	column := c.Params("id")

	result := initializers.DB.Unscoped().Delete(&models.Column{}, column)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
