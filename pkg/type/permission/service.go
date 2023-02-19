package permission

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// returns the list of all permissions as JSON.
// permission entries are filtered based on "collection" query parameter and the user id in the gin context.
func GetAllPermissions(c *gin.Context) {
	permissionsFilteredByCollection := GetPermissionsForCollection(c, permissions)
	permissionsFilteredByUserId := FilterEntriesBasedOnUserId(c, permissionsFilteredByCollection)

	c.IndentedJSON(http.StatusOK, permissionsFilteredByUserId)
}

// filters the list of all permissions based on the "collection" query parameter in the gin context;
// if the query parameter is not set, all permissions are returned
func GetPermissionsForCollection(c *gin.Context, permissionEntries []Permission) []Permission {
	collection := c.Query("collectionID")

	if collection == "" {
		return permissionEntries
	}

	filteredPermissions := []Permission{}
	for _, entry := range permissionEntries {
		if entry.Data.CollectionId == collection {
			filteredPermissions = append(filteredPermissions, entry)
		}
	}

	return filteredPermissions
}

func GetHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}