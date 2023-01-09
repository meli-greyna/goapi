package router

import (
	"github.com/gin-gonic/gin"
	"github.com/meli-greyna/goapi/handlers"
)

type Router struct {
	Server   *gin.Engine
	Handlers *handlers.Handlers
}

type resource struct {
	method   string
	path     string
	callback func(*gin.Context)
}

func (r *Router) MapPaths() {
	resources := []resource{
		{"GET", "/ping", r.Handlers.Ping},
		{"GET", "/products", r.Handlers.GetProducts},
		{"GET", "/products/:id", r.Handlers.GetProductById},
		{"GET", "/products/search", r.Handlers.SearchProducts},
		{"POST", "/products", r.Handlers.CreateProduct},
	}

	for _, resource := range resources {
		r.Server.Handle(resource.method, resource.path, resource.callback)
	}
}
