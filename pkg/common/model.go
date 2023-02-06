package common

import "time"

// identifier for types of things.
// e.g. group=herzog.fyi, type=postcard, schemaVersion=1
type TypeIdentifier struct {
	Group           string         `json:"id"`
	Type    string      `json:"type"`
	SchemaVersion    uint      `json:"schemaVersion"`
	
}

// metadata for every (sub) thing, heavily inspired by https://pkg.go.dev/golang.org/x/build/kubernetes/api#ObjectMeta
type ThingMetadata struct {
	UID UID `json:"uThingid,omitempty"`
	Name string `json:"name,omitempty"`
	Generation int64 `json:"generation,omitempty"`
	// reference to the thing that owns this (sub-) thing, i.e. the root thing
	// is only empty for root things - in this case the Collections list must NOT be empty
	OwnerThingRef string `json:"ownerThingRef,omitempty"`
	// reference to the direct parent (sub-) thing
	ParentRef string `json:"parentRef,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
	UpdateTimestamp time.Time `json:"updateTimestamp,omitempty"`
	// must be set for root things and reference at least one collection.
	// for sub-things, the collections list must be empty (only root things can be in a collection)
	Collections []UID `json:"collections,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// every (sub) thing has to have a UID but the UID format is up to the type (i.e. every microservice can choose its own UID implementation)
type UID string

