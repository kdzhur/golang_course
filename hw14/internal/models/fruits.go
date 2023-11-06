package models

import "fmt"

type Fruits struct {
	Fruits []Fruit
}

type Fruit struct {
	Name   string  `json:"name"`
	Size   float64 `json:"size"`
	Weight int     `json:"weight"`
}

func (f *Fruits) GetFruitByName(name string) *Fruit {
	for _, fruit := range f.Fruits {
		if fruit.Name == name {
			return &fruit
		}
	}
	return nil
}

func (f *Fruits) GetAllFruits() []Fruit {
	return f.Fruits
}

func (f *Fruits) AddFruit(fruit *Fruit) error {
	for _, f := range f.Fruits {
		if f.Name == fruit.Name {
			return fmt.Errorf("the fruit %v already exist", fruit.Name)
		}
	}
	f.Fruits = append(f.Fruits, *fruit)

	return nil
}

func (f *Fruits) DeleteFruit(fruit *Fruit) error {
	for i, fr := range f.Fruits {
		if fr.Name == fruit.Name {
			f.Fruits = append(f.Fruits[:i], f.Fruits[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("failed to delete fruit. Fruit %v not found", fruit.Name)
}
