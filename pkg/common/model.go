package common

import "time"

// identifier for types of things.
// e.g. group=herzog.fyi, type=postcard, schemaVersion=1
type TypeIdentifier struct {
	// a group is a domain name, e.g. herzog.fyi, optionally with a path, e.g. herzog.fyi/sammeles
	Group string `json:"id"`
	// the type of the thing, e.g. postcard.
	// should be short (often only one word) and not contain spaces or special characters besides _ and -
	Type string `json:"type"`
	// the schema version of the thing, e.g. 1.
	// the version is monotonically increasing and signals major changes to the type, possibly breaking backwards compatibility
	SchemaVersion int64 `json:"schemaVersion"`
}

// metadata for every (sub) thing, heavily inspired by https://pkg.go.dev/golang.org/x/build/kubernetes/api#ObjectMeta
type ThingMetadata struct {
	UID        UID    `json:"uThingid,omitempty"`
	Name       string `json:"name,omitempty"`
	Generation int64  `json:"generation,omitempty"`
	// reference to the direct parent (sub-) thing.
	// this is empty for root things.
	ParentRef         string    `json:"parentRef,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
	UpdateTimestamp   time.Time `json:"updateTimestamp,omitempty"`
	// collections that this (sub) thing belongs to.
	// for sub-things this will be synced eventually with the collections of the root thing (eventual consistency)
	Collections []UID             `json:"collections,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// every (sub) thing has to have a UID but the UID format is up to the type (i.e. every microservice can choose its own UID implementation)
type UID string
