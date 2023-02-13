package collection
import (
	"github.com/herzogf/sammeles/pkg/common"
)

func TypeIdentifier() common.TypeIdentifier {
	return common.TypeIdentifier{
		Group: "core.sammel.es",
		Type: "collection",
		SchemaVersion: 1,
	}
}

const TypePlural = "collections"