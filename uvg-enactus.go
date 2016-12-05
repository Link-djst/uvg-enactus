package main

import (
  "log"
  "os"
  "github.com/gin-gonic/gin"
)

func main() {
  port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set. Running localhost:8000")
    port = "8000"
	}

  r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H {
      "message": "pong",
    })
  })

  r.Run(":" + port)
}
