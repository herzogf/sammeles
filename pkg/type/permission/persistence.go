package permission

import (
	"time"

	"github.com/herzogf/sammeles/pkg/common"
)

var permissions = []Permission{
	{
		TypeIdentifier: TypeIdentifier(),
		Metadata: common.ThingMetadata{
			UID:               "collection-2-permission-1",
			Name:              "Sample Read Permission",
			Generation:        1,
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
			Collections: []common.UID{
				"collection-1",
			},
			Annotations: map[string]string{
				"color": "#00ff00",
			},
		},
		Data: PermissionData{
			UserId:       "user-1",
			CollectionId: "collection-2",
			Type:         Read,
		},
	},
	{
		TypeIdentifier: TypeIdentifier(),
		Metadata: common.ThingMetadata{
			UID:               "collection-2-permission-2",
			Name:              "Sample Write Permission",
			Generation:        1,
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
			Collections: []common.UID{
				"collection-1",
			},
			Annotations: map[string]string{
				"color": "#00ff00",
			},
		},
		Data: PermissionData{
			UserId:       "user-2",
			CollectionId: "collection-2",
			Type:         Write,
		},
	},
	{
		TypeIdentifier: TypeIdentifier(),
		Metadata: common.ThingMetadata{
			UID:               "collection-2-permission-3",
			Name:              "Sample Admin Permission",
			Generation:        1,
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
			Collections: []common.UID{
				"collection-1",
			},
			Annotations: map[string]string{
				"color": "#00ff00",
			},
		},
		Data: PermissionData{
			UserId:       "user-3",
			CollectionId: "collection-2",
			Type:         Admin,
		},
	},
}
