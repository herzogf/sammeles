package permission

import (
	"github.com/herzogf/sammeles/pkg/common"
)

type Permission struct {
	common.TypeIdentifier `json:",inline"`
	Metadata              common.ThingMetadata `json:"metadata,omitempty"`
	Data                  PermissionData       `json:"data,omitempty"`
}

type PermissionData struct {
	UserId     string         `json:"userId,omitempty"`
	CollectionId string         `json:"collectionId,omitempty"`
	Type         PermissionType `json:"type,omitempty"`
}

// enum for the permission type, either "write"" or "read"
type PermissionType string

const (
	// read for the collection
	Read  PermissionType = "read"
	// read + write for the collection
	Write PermissionType = "write"
	// write + create/delete permissions for the collection
	Admin PermissionType = "admin"
)
