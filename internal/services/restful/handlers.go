package restful

import (
	"net/http"

	"github.com/Onnywrite/lms-final/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func getStatus(c *gin.Context) {
	// faked
	c.JSON(http.StatusOK, gin.H{
		"id":     1567,
		"status": "calculating",
		"done":   "97.8%",
	})
}

func getServers(c *gin.Context) {

}

func postNew(c *gin.Context) {
	body := models.Expression{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "could not parse JSON body request",
		})
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}
