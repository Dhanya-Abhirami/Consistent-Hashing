package main

import (
	"server/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	docs "server/docs"
)

var hashRing *utils.HashRing
type IdInput struct {
	Id  string `json:"id" binding:"required"`
}
  
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

// @BasePath /api/v1

// Home godoc
// @Summary home
// @Schemes
// @Description home
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {string} home
// @Router / [get]
func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data":1})
}

// Add godoc
// @Summary	add
// @Schemes
// @Description add
// @Tags add
// @Accept json
// @Produce json
// @Param account body IdInput true "Add node"
// @Success 200 {string} add
// @Router /add [post]
func add(c *gin.Context) {
	var input IdInput
	if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }
	hashRing.AddNode(input.Id)
	c.JSON(http.StatusOK, gin.H{"id":input.Id})
}

// Remove godoc
// @Summary remove
// @Schemes
// @Description remove
// @Tags remove
// @Accept json
// @Produce json
// @Param account body IdInput true "Remove node"
// @Success 200 {string} remove
// @Router /remove [delete]
func remove(c *gin.Context) {
	var input IdInput
	if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }
	hashRing.RemoveNode(input.Id)
	c.JSON(http.StatusOK, gin.H{"id":input.Id})
}

func main() {
	hashRing = utils.NewHashRing()
	router := gin.Default()
	router.Use(CORS) 
	docs.SwaggerInfo.BasePath = "/api/v1"
   	v1 := router.Group("/api/v1")
	{
		v1.GET("/",home)
		v1.POST("/add", add)
		v1.DELETE("/remove", remove)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
