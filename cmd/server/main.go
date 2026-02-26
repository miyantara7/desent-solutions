package main

import (
	"log"

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

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
