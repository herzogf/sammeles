package permission

import (
	"time"

	"github.com/herzogf/sammeles/pkg/common"
)

var permissions = []Permission{
	{
		TypeIdentifier: TypeIdentifier(),
		Metadata: common.ThingMetadata{
			UID:               "abc123",
			Name:              "permission-1",
			Generation:        1,
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
			Collections: []common.UID{
				"collection1",
				"collection2",
			},
			Annotations: map[string]string{
				"color": "#00ff00",
			},
		},
		Data: PermissionData{
			UserId:       "user1",
			CollectionId: "collection1",
			Type:         Read,
		},
	},
}
