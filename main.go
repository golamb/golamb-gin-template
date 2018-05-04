package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "gin lambda server.",
	})
}

func routerEngine() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	mux := gin.New()
	mux.GET("/", rootHandler)

	return mux
}

func main() {
	addr := ":" + os.Getenv("PORT")
	log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
}
