package main

import (
	"fmt"
	"log"

	"github.com/meli-greyna/goapi/product"
)

func main() {
	products, err := product.Ingest("products.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
}
