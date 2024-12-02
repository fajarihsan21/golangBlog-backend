package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := "8080"
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ReadTimeout:  10 * time.Minute, // Timeout for reading the request
		WriteTimeout: 10 * time.Minute, // Timeout for writing the response
		IdleTimeout:  10 * time.Minute, // Timeout for idle connections
	}

	fmt.Println(fmt.Sprintf("LINTAS SR TOOLS SERVICE STARTED ON PORT %d", port))

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
	router.Run(fmt.Sprintf(":%v", port))
}
