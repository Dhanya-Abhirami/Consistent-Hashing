package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"server/handlers"
	// swaggerfiles "github.com/swaggo/files" // swagger embed files
	// docs "server/docs"
)

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

var router *gin.Engine

func main() {
	router := gin.Default()
	router.Use(CORS) 
	v1 := router.Group("/api/v1")
	{
		v1.POST("/hashring", handlers.HashRing)
		v1.PUT("/server/:id", handlers.Add)
		v1.DELETE("/server/:id", handlers.Remove)
		v1.GET("/mapping/:id", handlers.Mapping)
		v1.GET("/mapping/all", handlers.MappingAll)
	}
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
	
}
