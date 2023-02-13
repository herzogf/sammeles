package collection

import (
	"github.com/herzogf/sammeles/pkg/common"
	"time"
)

var collections = []Collection{
	{
		TypeIdentifier: TypeIdentifier(),
		Metadata: common.ThingMetadata{
			UID:               "collection-1",
			Name:              "Global Permission Collection",
			Generation:        1,
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
			// currently collections are typically NOT contained in other collections
			Collections: []common.UID{},
		},
		Data: CollectionData{
			Things: []common.ThingIdentifier{
				{
					Group: "core.sammel.es",
					Type:  "permission",
					UID:   "collection-2-permission-1",
				},
				{
					Group: "core.sammel.es",
					Type:  "permission",
					UID:   "collection-2-permission-2",
				},
				{
					Group: "core.sammel.es",
					Type:  "permission",
					UID:   "collection-2-permission-3",
				},
			},
		},
	},
	{
		TypeIdentifier: TypeIdentifier(),
		Metadata: common.ThingMetadata{
			UID:               "collection-2",
			Name:              "Just a test collection",
			Generation:        1,
			CreationTimestamp: time.Now(),
			UpdateTimestamp:   time.Now(),
			Collections: []common.UID{},
		},
		Data: CollectionData{
			Things: []common.ThingIdentifier{
				{
					Group: "sammeles.herzog.fyi",
					Type:  "postcard",
					UID:   "postcard-1",
				},
				{
					Group: "sammeles.herzog.fyi",
					Type:  "postcard",
					UID:   "postcard-2",
				},
			},
		},
	},
}
