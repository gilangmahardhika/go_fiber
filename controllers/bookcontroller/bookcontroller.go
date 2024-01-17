package bookcontroller

import (
	"github.com/gilangmahardhika/golang-web/models"
	"github.com/gilangmahardhika/golang-web/responses"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book

	models.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": books})
}

func Show(c *fiber.Ctx) error {
	return nil
}

func Create(c *fiber.Ctx) error {
	var book models.Book
	var response responses.Error
	if err := c.BodyParser(&book); err != nil {
		response.Message = err
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": response,
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		response.Message = err
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": response,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": book,
	})
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
