package collection

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// returns the list of all collections as JSON
func GetAllCollections(c *gin.Context) {
	filteredCollectionsBasedOnUserId, err := FilterEntriesBasedOnUserId(c, collections)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, filteredCollectionsBasedOnUserId)
	}
}

func GetHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
