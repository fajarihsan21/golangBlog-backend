package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		errEnv := godotenv.Load()
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("APPPORT")
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ReadTimeout:  10 * time.Minute, // Timeout for reading the request
		WriteTimeout: 10 * time.Minute, // Timeout for writing the response
		IdleTimeout:  10 * time.Minute, // Timeout for idle connections
	}

	fmt.Println(fmt.Sprintf("SERVICE STARTED ON PORT %d", port))

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
	router.Run(fmt.Sprintf(":%v", port))
}
