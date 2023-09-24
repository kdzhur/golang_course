// 1. Створити інтерфейс «Посилка» й реалізувати його для двох класів — «Коробка» і «Конверт».
// Для кожної поштової посилки необхідно зберігати адресу отримувача й відправника.
// Додати сортувальний відділ, який залежно від типу посилки відправляє її тим або іншим шляхом.

// 2. Створити інтерфейс «Транспортний засіб» і реалізувати його для класів «Автомобіль», «Потяг», «Літак».
// Кожен із цих транспортних засобів має мати методи «Рухатися», «Зупинятися» та «Змінювати швидкість».
// Додатково створити інтерфейс «Транспортний засіб» і реалізувати його для класів «Автомобіль», «Потяг», «Літак»,
//  які мають методи «Приймати пасажирів» та «Висаджувати пасажирів». Укінці створити клас «Маршрут», який містить список транспортних засобів,
//  які пройшли по заданому маршруту. Клас «Маршрут» має мати методи «Додавати транспортний засіб до маршруту» та «Показувати список транспортних засобів на маршруті».
//  Тепер цей маршрут мусить пройти ваш подорожувальник із виводом його подорожі на екран.

// Файли зберігати в різних пакетах.

package main

import (
	"fmt"
	"my_interfaces/internal/parcel"
	"my_interfaces/internal/route"
	"my_interfaces/internal/vehicle"
)

func main() {
	fmt.Print("1.===================\n\n")

	s := parcel.Sender{
		Customer: parcel.Customer{
			Name:    "Sherlock",
			Address: "London, Baker str. 211b",
		},
	}

	r := parcel.Recipient{
		Customer: parcel.Customer{
			Name:    "Watson",
			Address: "London, Baker str. 311a",
		},
	}

	envelope := parcel.Envelope{
		Parcel: parcel.Parcel{
			Name:      "An envelope to Watson",
			Sender:    &s,
			Recipient: &r,
		},
	}

	box := parcel.Box{
		Parcel: parcel.Parcel{
			Name:      "A box to Watson",
			Sender:    &s,
			Recipient: &r,
		},
	}

	parcel.SendParcelTo(&envelope)
	parcel.SendParcelTo(&box)

	fmt.Print("2.===================\n\n")

	car := vehicle.Car{
		Vehicle: vehicle.Vehicle{
			Name:  "Car",
			Speed: 0,
		},
		Brand: "BMW",
	}

	train := vehicle.Train{
		Vehicle: vehicle.Vehicle{
			Name:  "Train",
			Speed: 0,
		},
	}

	airplane := vehicle.Train{
		Vehicle: vehicle.Vehicle{
			Name:  "Airplane",
			Speed: 0,
		},
	}

	route := route.Route{
		Name: "Kyiv - Warsaw",
	}

	car.ChangeSpeed(50)
	car.Move()
	car.Stop()
	car.Move()

	car.TakePassenger(4)
	car.DropPassenger(1)

	route.AddVehicle(&car)
	route.AddVehicle(&train)
	route.AddVehicle(&airplane)
	route.PrintVehicles()
}
