package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleHello(c *gin.Context) {
	log.Println("DEBUG: handle hello started")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world",
	})
	log.Println("DEBUG: handle hello finished")
}
