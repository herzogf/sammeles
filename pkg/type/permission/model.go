package permissions

import (
	"github.com/herzogf/sammeles/pkg/common"
)

type Permission struct {
	common.TypeIdentifier `json:",inline"`
	Metadata              common.ThingMetadata `json:"metadata,omitempty"`
	Data                  PermissionData       `json:"data,omitempty"`
}

type PermissionData struct {
	Username     string         `json:"username,omitempty"`
	CollectionId string         `json:"collectionId,omitempty"`
	Type         PermissionType `json:"type,omitempty"`
}

// enum for the permission type, either "write"" or "read"
type PermissionType string

const (
	Read  PermissionType = "read"
	Write PermissionType = "write"
)
