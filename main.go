package main

import (
	"github.com/gin-gonic/gin"
	"live-config/server/handler"
)

func main() {
	initHttpServer()
}

func initHttpServer() {
	r := gin.Default()

	c := handler.New()

	group := r.Group("/property")

	group.GET("/:application/:profile/:label", c.GetByApplicationProfileAndLabel)
	group.POST("/", c.Create)
	group.PUT("/:id", c.Modify)

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
