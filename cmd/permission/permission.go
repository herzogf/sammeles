package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/herzogf/sammeles/pkg/type/permission"
)

// main function for the permission type / service
func main() {
	router := gin.Default()

	typeIdentifier := permission.TypeIdentifier()
	routes := router.Group("/api/types/" + typeIdentifier.Group)
	routes.GET("/" + typeIdentifier.Type + "/entries", permission.GetAllPermissions)
	routes.GET("/" + permission.TypePlural + "/entries", permission.GetAllPermissions)

	router.Run(":8080")
	fmt.Println("Serving type 'permission' on port 8080")
}
