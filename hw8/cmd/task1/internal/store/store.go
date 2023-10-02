package store

import (
	"concurrency2/cmd/task1/internal/product"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Store struct {
	Title    string
	Products []*product.Product
}

func NewStore(title string) *Store {
	return &Store{
		Title:    title,
		Products: []*product.Product{},
	}
}

func (s *Store) AddToStore(p *product.Product) {
	s.Products = append(s.Products, p)
}

func (s *Store) FillStore(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(byteValue, &s.Products)
}

func (s *Store) ShowStore() {
	for _, product := range s.Products {
		fmt.Printf("======\nName: %v\nModel: %v\nManufacturer: %v\nPrice: %v\n", product.Name, product.Model, product.Manufacturer, product.Price)
	}
}
