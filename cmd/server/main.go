package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/miyantara7/desent-solutions/internal/container"
)

func main() {
	ctn, err := container.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer ctn.Delete()

	r := ctn.Get(container.ContainerRouter).(*gin.Engine)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
