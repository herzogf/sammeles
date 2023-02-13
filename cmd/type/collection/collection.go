package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/herzogf/sammeles/pkg/common"
	"github.com/herzogf/sammeles/pkg/type/collection"
)

// main function for the collection type / service
func main() {
	router := gin.Default()

	typeIdentifier := collection.TypeIdentifier()
	routes := router.Group("/api/types/" + typeIdentifier.Group)
	routes.GET("/" + typeIdentifier.Type + "/entries", collection.GetAllCollections)
	routes.GET("/" + collection.TypePlural + "/entries", collection.GetAllCollections)

	router.GET("/health", collection.GetHealth)

	port := getPort()
	fmt.Printf("Serving type 'collection' on port %d\n", port)

	// graceful shutdown for gin, see https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()


	common.RegisterOneAndOnlyType(collection.TypeIdentifier(), port)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	common.DeregisterOneAndOnlyType(collection.TypeIdentifier())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("Shutdown finished")
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