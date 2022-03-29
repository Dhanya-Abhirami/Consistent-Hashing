package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Home godoc
// @Summary mapping
// @Schemes
// @Description mapping
// @Tags mapping
// @Accept json
// @Param id path string true "Server ID"
// @Produce json
// @Success 200 {string} mapping
// @Failure 405 {string} mapping
// @Router /mapping/{id} [get]
func Mapping(c *gin.Context) {
	if models.HASHRING!=nil && models.HASHRING.Servers!=nil{
		id := c.Param("id")
		server := models.HASHRING.GetMapping(id)
		c.JSON(http.StatusOK, server)
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
// @Failure 405 {string} mappingAll
// @Router /mapping/all [get]
func MappingAll(c *gin.Context) {
	if models.HASHRING!=nil && models.HASHRING.Servers!=nil{
		c.JSON(http.StatusOK, models.HASHRING)
	} else{             
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}