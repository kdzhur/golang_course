package models

type Fruits struct {
	Fruits []Fruit
}

type Fruit struct {
	Name   string
	Size   float64
	Weight int
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
