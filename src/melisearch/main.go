package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryparra/melisearch/src/melisearch/controllers/myml"
	"github.com/mauryparra/melisearch/src/melisearch/controllers/ping"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", ping.Ping)
	router.GET("/myml/:userID", myml.GetInfo)
	router.Run(port)
}
