package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/meli-greyna/goapi/handlers"
	"github.com/meli-greyna/goapi/product"
	"github.com/meli-greyna/goapi/router"
)

func main() {
	products, err := product.Ingest("products.json")
	if err != nil {
		log.Fatal(err)
	}

	router := router.Router{
		Server:   gin.Default(),
		Handlers: &handlers.Handlers{Products: &products},
	}

	router.MapPaths()

	router.Server.Run(":8080")
}
