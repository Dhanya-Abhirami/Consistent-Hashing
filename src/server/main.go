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
type keysInput struct {
	Keys  int `json:"keys" binding:"required"`
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

// Add godoc
// @Summary	keysInput
// @Schemes
// @Description keysInput
// @Tags keysInput
// @Accept json
// @Produce json
// @Param account body keysInput true "Create hashring"
// @Success 200 {string} hashring
// @Router /hashring [post]
func hashring(c *gin.Context) {
	var input keysInput
	err := c.BindJSON(&input);
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else{
		hashRing = utils.NewHashRing(input.Keys)
		c.JSON(http.StatusOK, "Created Hash Ring")
	}
}

// Add godoc
// @Summary	add
// @Schemes
// @Description add
// @Tags add
// @Accept json
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {string} add
// @Failure 405 {string} add
// @Router /node/{id} [put]
func add(c *gin.Context) {
	if hashRing!=nil && hashRing.Nodes!=nil{
		id := c.Param("id")
		remap := hashRing.AddNode(id)
		c.JSON(http.StatusOK, gin.H{"remap":remap})
	} else{
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

// Remove godoc
// @Summary remove
// @Schemes
// @Description remove
// @Tags remove
// @Accept json
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {string} remove
// @Failure 404 {string} remove
// @Failure 405 {string} remove
// @Router /node/{id} [delete]
func remove(c *gin.Context) {
	if hashRing!=nil && hashRing.Nodes!=nil {
		id := c.Param("id")
		remap,err := hashRing.RemoveNode(id)
		if err!=nil {
			c.JSON(http.StatusNotFound, err.Error())
		} else{
			c.JSON(http.StatusOK, gin.H{"remap":remap})
		}
		
	} else{             
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
	
}

// Home godoc
// @Summary mapping
// @Schemes
// @Description mapping
// @Tags mapping
// @Accept json
// @Param id path string true "Node ID"
// @Produce json
// @Success 200 {string} mapping
// @Failure 404 {string} mapping
// @Router /mapping/{id} [get]
func mapping(c *gin.Context) {
	if hashRing!=nil && hashRing.Nodes!=nil{
		id := c.Param("id")
		server,err := hashRing.GetMapping(id)
		if err!=nil {
			c.JSON(http.StatusNotFound, err.Error())
		} else{
			c.JSON(http.StatusOK, server)
		}
		
	} else{             
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
	
}

// Home godoc
// @Summary mappingAll
// @Schemes
// @Description mappingAll
// @Tags mappingAll
// @Accept json
// @Produce json
// @Success 200 {string} mappingAll
// @Failure 404 {string} mappingAll
// @Router /mapping/all [get]
func mappingAll(c *gin.Context) {
	if hashRing!=nil && hashRing.Nodes!=nil{
		c.JSON(http.StatusOK, hashRing.Nodes)
	} else{             
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

func main() {
	router := gin.Default()
	router.Use(CORS) 
	docs.SwaggerInfo.BasePath = "/api/v1"
   	v1 := router.Group("/api/v1")
	{
		v1.POST("/hashring", hashring)
		v1.PUT("/node/:id", add)
		v1.DELETE("/node/:id", remove)
		v1.GET("/mapping/:id", mapping)
		v1.GET("/mapping/all", mappingAll)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
