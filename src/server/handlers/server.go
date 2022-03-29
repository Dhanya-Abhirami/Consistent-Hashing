package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Add godoc
// @Summary	add
// @Schemes
// @Description add
// @Tags add
// @Accept json
// @Produce json
// @Param id path string true "Server ID"
// @Success 200 {string} add
// @Failure 405 {string} add
// @Router /server/{id} [put]
func Add(c *gin.Context) {
	if models.HASHRING!=nil && models.HASHRING.Servers!=nil{
		id := c.Param("id")
		remap := models.HASHRING.AddServer(id)
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
// @Param id path string true "Server ID"
// @Success 200 {string} remove
// @Failure 404 {string} remove
// @Failure 405 {string} remove
// @Router /server/{id} [delete]
func Remove(c *gin.Context) {
	if models.HASHRING!=nil && models.HASHRING.Servers!=nil {
		id := c.Param("id")
		remap,err := models.HASHRING.RemoveServer(id)
		if err!=nil {
			c.JSON(http.StatusNotFound, err.Error())
		} else{
			c.JSON(http.StatusOK, gin.H{"remap":remap})
		}
		
	} else{             
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
	
}

