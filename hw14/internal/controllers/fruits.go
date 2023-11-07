package controllers

import (
	"errors"
	"fmt"
	"mvc/internal/models"
	"mvc/internal/views"

	"github.com/gofiber/fiber/v2"
)

func ListAllFruits(c *fiber.Ctx) error {
	if fruitsList, ok := c.Locals("fruits").(*models.Fruits); ok {
		return views.JsonFruits(c, (fruitsList.GetAllFruits()))
	}
	return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to list fruits"))
}

func GetFruit(c *fiber.Ctx) error {
	fruit := c.Query("name")

	if fruitsList, ok := c.Locals("fruits").(*models.Fruits); ok {
		return views.JsonFruits(c, fruitsList.GetFruitByName(fruit))
	}
	return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to get fruit"))
}

func AddFruit(c *fiber.Ctx) error {
	fruit := new(models.Fruit)

	if err := c.BodyParser(&fruit); err != nil {
		return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to add fruit"), err)
	}

	if fruitsList, ok := c.Locals("fruits").(*models.Fruits); ok {
		if err := fruitsList.AddFruit(fruit); err != nil {
			return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to add fruit"), err)
		}
		return views.JsonFruits(c, fruit)
	}
	return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to add fruit"))
}

func DeleteFruit(c *fiber.Ctx) error {
	fruitName := c.Params("name")

	fruit := models.Fruit{
		Name: fruitName,
	}

	if fruitsList, ok := c.Locals("fruits").(*models.Fruits); ok {
		if err := fruitsList.DeleteFruit(&fruit); err != nil {
			return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to delete fruit"), err)
		}
		return views.JsonFruits(c, fruit)
	}
	return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to delete fruit"))
}
