package controllers

import (
	"errors"
	"fmt"
	"mvc/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ListAllFruits(c *fiber.Ctx) error {
	if fruitsList, ok := c.Locals("fruits").(*models.Fruits); ok {
		c.JSON(fruitsList.GetAllFruits())
		return nil
	}
	return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to list fruits"))
}

func GetFruit(c *fiber.Ctx) error {
	fruit := c.Query("name")

	if fruitsList, ok := c.Locals("fruits").(*models.Fruits); ok {
		c.JSON(fruitsList.GetFruitByName(fruit))
		return nil
	}
	return errors.Join(fiber.ErrInternalServerError, fmt.Errorf("failed to get fruit"))
}
