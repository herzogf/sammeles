package collection

import (
	"github.com/herzogf/sammeles/pkg/common"
)

// Collections contain concrete things
type Collection struct {
	common.TypeIdentifier `json:",inline"`
	Metadata              common.ThingMetadata `json:"metadata,omitempty"`
	Data                  CollectionData       `json:"data,omitempty"`
}

// CollectionData is the data of a collection
type CollectionData struct {
	// Things is a list of top-level things that are contained in the collection.
	// This list does NOT contain sub-things (those can be traversed via the top-level things).
	Things     []common.ThingIdentifier         `json:"userId,omitempty"`
}