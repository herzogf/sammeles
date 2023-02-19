package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/herzogf/sammeles/pkg/common"
)

// filters the given permission list based on the user id in the gin context.
// admin and internal users get all permissions, anonymous users get no permissions,
// and all other users get only the permissions that reference their user id.
func FilterEntriesBasedOnUserId(c *gin.Context, permissionEntries []Permission) []Permission {
	userId := common.GetUserIdFromGinContext(c)

	switch userId {
	case common.USER_ID_ADMIN, common.USER_ID_INTERNAL:
		return permissionEntries
	case "":
		return []Permission{}
	default:
		filteredEntries := []Permission{}
		for _, entry := range permissionEntries {
			if entry.Data.UserId == userId {
				filteredEntries = append(filteredEntries, entry)
			}
		}
	
		return filteredEntries
	}
}
