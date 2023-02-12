package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herzogf/sammeles/pkg/common"
	"github.com/herzogf/sammeles/pkg/type/permission"
)

// main function for the permission type / service
func main() {
	router := gin.Default()

	typeIdentifier := permission.TypeIdentifier()
	routes := router.Group("/api/types/" + typeIdentifier.Group)
	routes.GET("/" + typeIdentifier.Type + "/entries", permission.GetAllPermissions)
	routes.GET("/" + permission.TypePlural + "/entries", permission.GetAllPermissions)

	router.GET("/health", permission.GetHealth)

	port := getPort()
	common.RegisterOneAndOnlyType(permission.TypeIdentifier(), port)

	router.Run(fmt.Sprintf(":%d", port))
	fmt.Printf("Serving type 'permission' on port %d\n", port)
}

func getPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}
	return p
}