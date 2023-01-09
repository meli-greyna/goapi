package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meli-greyna/goapi/product"
)

type Handlers struct {
	Products *[]product.Product
}

func (h *Handlers) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (h *Handlers) GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, *h.Products)
}

func (h *Handlers) GetProductById(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var prod product.Product

	for _, p := range *h.Products {
		if p.Id == int(id) {
			prod = p
		}
	}

	if prod == (product.Product{}) {
		ctx.String(http.StatusNotFound, "product not found")
	} else {
		ctx.JSON(http.StatusOK, prod)
	}
}

func (h *Handlers) SearchProducts(ctx *gin.Context) {
	priceQuery, ok := ctx.GetQuery("priceGt")
	if !ok {
		ctx.JSON(http.StatusOK, h.Products)
	} else {
		price, _ := strconv.ParseFloat(priceQuery, 64)
		var products []product.Product

		for _, p := range *h.Products {
			if p.Price >= price {
				products = append(products, p)
			}
		}

		ctx.JSON(http.StatusOK, products)
	}
}
