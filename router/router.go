package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meli-greyna/goapi/product"
)

type Router struct {
	Server    *gin.Engine
	Products  *[]product.Product
	resources []resource
}

type resource struct {
	method   string
	path     string
	callback func(*gin.Context)
}

func (r *Router) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (r *Router) GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, *r.Products)
}

func (r *Router) GetProductById(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var prod product.Product

	for _, p := range *r.Products {
		if p.Id == int(id) {
			prod = p
		}
	}

	ctx.JSON(http.StatusOK, prod)
}

func (r *Router) GetProductsOfPrice(ctx *gin.Context) {
	price, _ := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	var products []product.Product

	for _, p := range *r.Products {
		if p.Price >= price {
			products = append(products, p)
		}
	}

	ctx.JSON(http.StatusOK, products)
}

func (r *Router) MapPaths() {
	r.resources = []resource{
		{"GET", "/ping", r.Ping},
		{"GET", "/products", r.GetProducts},
		{"GET", "/products/:id", r.GetProductById},
		{"GET", "/products/search", r.GetProductsOfPrice},
	}

	for _, resource := range r.resources {
		r.Server.Handle(resource.method, resource.path, resource.callback)
	}
}
