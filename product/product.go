package product

import (
	"encoding/json"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func Ingest(filename string) ([]Product, error) {
	var products = []Product{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &products)

	return products, nil
}
