package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yendelevium/sebastian/internal/routes"
)

func main() {
	log.Println("sebastian: at your service")
	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// Suggested by gin official docs
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	routes.AddRoutes(&r.RouterGroup)

	// Start server on port 8080
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
