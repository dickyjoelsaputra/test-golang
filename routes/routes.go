package routes

import (
	"test-golang/controllers"
	"test-golang/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin init
	r := gin.Default()
	r.Use(static.Serve("/Images", static.LocalFile("./Images", true)))

	v1 := r.Group("/api")
	{
		// TEST 1 - 5
		v1.POST("/test1", controllers.TestSatu)
		v1.POST("/test2", controllers.TestDua)
		v1.GET("/test3", controllers.TestTiga)
		v1.POST("/test4", controllers.TestEmpat)
		v1.POST("/test5", controllers.TestLima)

		// AUTH
		v1.POST("register", controllers.Register)
		v1.POST("login", controllers.Login)
		// PRODUCT
		v1.GET("products/:page", middleware.Auth, controllers.ProductIndex)
		v1.GET("product/:id", middleware.Auth, controllers.ProductShow)
		v1.POST("product", middleware.Auth, controllers.ProductCreate)
		v1.PUT("product/:id", middleware.Auth, controllers.ProductUpdate)
		v1.DELETE("product/:id", middleware.Auth, controllers.ProductDelete)
		// GLOBAL IMAGE
		r.GET("images/:image", func(c *gin.Context) {
			namaGambar := c.Param("image")
			c.File("./images/" + namaGambar)
		})
	}

	return r
}
