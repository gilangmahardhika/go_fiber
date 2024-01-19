package bookcontroller

import (
	"github.com/gilangmahardhika/golang-web/models"
	"github.com/gilangmahardhika/golang-web/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book

	models.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": books})
}

func Show(c *fiber.Ctx) error {
	var book models.Book
	var response responses.Error

	id := c.Params("id")
	result := models.DB.First(&book, id)
	log.Info(book)

	if result.RowsAffected == 0 {
		log.Warn("Not Found")
		response.Message = "book not found"
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data": response,
		})
	}
	log.Info(book)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": book})
}

func Create(c *fiber.Ctx) error {
	var book models.Book
	var response responses.Error
	if err := c.BodyParser(&book); err != nil {
		log.Warn(err)
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": response,
		})

	}

	if err := models.DB.Create(&book).Error; err != nil {
		log.Warn(err)
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": response,
		})

	}
	log.Info("create book", book)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": book,
	})
}

func Update(c *fiber.Ctx) error {
	var book models.Book
	var response responses.Error

	id := c.Params("id")
	result := models.DB.First(&book, id)
	log.Info(book)

	if result.RowsAffected == 0 {
		log.Warn("Not Found")
		response.Message = "book not found"
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data": response,
		})
	}

	if err := c.BodyParser(&book); err != nil {
		log.Warn(err)
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": response,
		})

	}

	if err := models.DB.Save(&book).Error; err != nil {
		log.Warn(err)
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": response,
		})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": book})
}

func Delete(c *fiber.Ctx) error {
	var book models.Book
	var response responses.Error

	id := c.Params("id")
	result := models.DB.First(&book, id).Delete(&book)
	log.Info(book)

	if result.RowsAffected == 0 {
		log.Warn("Not Found")
		response.Message = "book not found"
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data": response,
		})
	}

	response.Message = "book has been deleted"
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response})
}
