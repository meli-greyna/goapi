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

func (h *Handlers) CreateProduct(ctx *gin.Context) {
	var req gin.H

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
		return
	}

	name, quantity, code_value, is_published, expiration, price := req["name"], req["quantity"], req["code_value"], req["is_published"], req["expiration"], req["price"]

	if is_published == nil {
		is_published = false
	}

	if name == nil ||
		quantity == nil ||
		code_value == nil ||
		expiration == nil ||
		price == nil {
		ctx.String(http.StatusBadRequest, "invalid product data")
		return
	}

	for _, p := range *h.Products {
		if p.CodeValue == req["code_value"] {
			ctx.String(http.StatusBadRequest, "duplicated code_value")
			return
		}
	}

	product := product.Product{
		Id:          (*h.Products)[len(*h.Products)-1].Id + 1,
		Name:        name.(string),
		Quantity:    int(quantity.(float64)),
		CodeValue:   code_value.(string),
		IsPublished: is_published.(bool),
		Expiration:  expiration.(string),
		Price:       price.(float64),
	}

	*h.Products = append(*h.Products, product)

	ctx.JSON(http.StatusOK, product)
}
