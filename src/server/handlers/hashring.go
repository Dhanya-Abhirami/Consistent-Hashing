package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type keysInput struct {
	Keys  int `json:"keys" binding:"required"`
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
func HashRing(c *gin.Context) {
	var input keysInput
	err := c.BindJSON(&input);
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else{
		models.HASHRING = models.NewHashRing(input.Keys)
		c.JSON(http.StatusOK, "Created Hash Ring")
	}
}

