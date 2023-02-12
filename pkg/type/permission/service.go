package permission

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// returns the list of all permissions as JSON
func GetAllPermissions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, permissions)
}

func GetHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}