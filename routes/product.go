package routes

import (
	"backend_go/constants"
	middleware "backend_go/middlewares"
	"backend_go/permissions"
	"backend_go/services"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.Use(middleware.AuthMiddlewareFlexible())
		products.GET("/", services.GetProducts)
		products.GET("/:id", services.GetProductByID)
		products.POST("/", permissions.RequirePermissions(constants.PermissionProductsCreate), services.CreateProduct)
		products.PATCH("/:id", permissions.RequirePermissions(constants.PermissionProductsUpdate), services.UpdateProduct)
		products.DELETE("/:id", permissions.RequirePermissions(constants.PermissionProductsDelete), services.DeleteProduct)
	}
}
