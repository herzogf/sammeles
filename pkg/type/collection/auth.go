package collection

import (
	"log"

	"github.com/herzogf/sammeles/pkg/common"
	"github.com/herzogf/sammeles/pkg/type/permission"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// Filter collections based on the user id in the gin context.
func FilterEntriesBasedOnUserId(c *gin.Context, collectionEntries []Collection) ([]Collection, error) {
	userId := common.GetUserIdFromGinContext(c)

	switch userId {
	case common.USER_ID_ADMIN, common.USER_ID_INTERNAL:
		return collectionEntries, nil
	case "":
		return []Collection{}, nil
	default:
		allowedCollectionIds, err := getCollectionIdsReadableByUser(userId)

		if err != nil {
			return  nil, err
		}

		filteredEntries := []Collection{}
		for _, entry := range collectionEntries {
			if slices.Contains(allowedCollectionIds, entry.Metadata.UID) {
				filteredEntries = append(filteredEntries, entry)
			}
		}

		return filteredEntries, nil
	}
}

// Get collections for the user id in the gin context.
func getCollectionIdsReadableByUser(userId string) ([]common.UID, error) {
	permissionTypeServiceUrl, err := common.FindTypeServiceUrl(permission.TypeIdentifier())
	log.Println("Discovered service url: " + permissionTypeServiceUrl)

	if err != nil {
		return nil, err
	}

	return []common.UID{
		"collection-2", // postcard collection
	}, nil
}
