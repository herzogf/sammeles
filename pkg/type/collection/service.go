package collection

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// returns the list of all collections as JSON
func GetAllCollections(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, collections)
}

func GetHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}