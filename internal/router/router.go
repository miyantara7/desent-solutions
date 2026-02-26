package router

import "github.com/gin-gonic/gin"

type Registrar interface {
	Register(r *gin.Engine)
}

func SetupRouter(registrars ...Registrar) *gin.Engine {
	r := gin.Default()

	for _, reg := range registrars {
		reg.Register(r)
	}

	return r
}
