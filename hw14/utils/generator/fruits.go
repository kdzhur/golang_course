package generator

import (
	"math/rand"
	"mvc/internal/models"
)

func GenerateRandomFruits(randomFruits []string) []models.Fruit {
	var fruits []models.Fruit

	for _, f := range randomFruits {
		fruit := models.Fruit{
			Name:   f,
			Size:   rand.Float64() * 10,
			Weight: rand.Intn(500),
		}
		fruits = append(fruits, fruit)
	}

	return fruits
}
