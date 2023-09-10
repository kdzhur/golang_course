// Розробити програму «Зоопарк». Завдання: 5 чи більше звірів повтікали, наглядач повинен їх зібрати. Кожну сутність (наглядач, звір, клітка тощо) представляти окремою структурою (zookeeper, animal, cage). Користуємось ембдінгом і методами. 2. Зареєструватися на https://github.com/. 3. Створити репозиторій і надіслати виконану домашню роботу на github.

package main

import (
	"fmt"
)

const (
	Snowy    string = "snowy"
	Aquatic  string = "aquatic"
	Savannah string = "savannah"
	Jungle   string = "jungle"
)

type animal struct {
	name         string
	habitatBiome string
}

func newAnimal(name, habitatBiome string) animal {
	a := animal{
		name:         name,
		habitatBiome: habitatBiome,
	}

	return a
}

type zookeeper struct {
	name  string
	cages []*cage // cages that zookeeper owns
}

func (z *zookeeper) cageAnimal(a animal) {
	for i := range z.cages {
		if z.cages[i].canFitAnimal(a) {
			z.cages[i].fitAnimal(a)
			fmt.Printf("Zookeper %v has caged animal %v to the cage #%v\n", z.name, a.name, z.cages[i].number)
			break
		}
		if i+1 == len(z.cages) {
			fmt.Printf("Zookeper %v has no free cage with %v biome for animal %v\n", z.name, a.habitatBiome, a.name)
		}
	}
}

type cage struct {
	number int
	biome  string
	animal animal
}

func (c *cage) fitAnimal(a animal) {
	c.animal = a
}

func (c *cage) canFitAnimal(a animal) bool {
	return (c.animal == animal{} && a.habitatBiome == c.biome)
}

func main() {

	animal1 := newAnimal("Lion", Savannah)
	animal2 := newAnimal("Penguin", Snowy)
	animal3 := newAnimal("Hippopotamus", Aquatic)
	animal4 := newAnimal("Arctic_bear", Snowy)
	animal5 := newAnimal("Monkey", Jungle)
	animal6 := newAnimal("Snake", Jungle)
	animal7 := newAnimal("Panther", Jungle)

	cages := []cage{
		{
			number: 1,
			biome:  Snowy,
		},
		{
			number: 2,
			biome:  Savannah,
			animal: animal1, // This animal didn't run away
		},
		{
			number: 3,
			biome:  Snowy,
		},
		{
			number: 4,
			biome:  Aquatic,
		},
		{
			number: 5,
			biome:  Jungle,
		},
		{
			number: 6,
			biome:  Jungle,
		},
	}

	z1 := zookeeper{
		name:  "Oleksiy",
		cages: []*cage{&cages[2], &cages[3], &cages[4], &cages[5]},
	}

	z2 := zookeeper{
		name:  "Kyrylo",
		cages: []*cage{&cages[0], &cages[1], &cages[2], &cages[3]},
	}

	// Catch animals
	z1.cageAnimal(animal2)
	z2.cageAnimal(animal3)
	z2.cageAnimal(animal4)
	z1.cageAnimal(animal5)
	z1.cageAnimal(animal6)

	// Oleksiy has no free room left. All cages are taken
	z1.cageAnimal(animal7)
	// Aswell as Kyrylo
	z2.cageAnimal(animal7)

	fmt.Println(cages)
}
