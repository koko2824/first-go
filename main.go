package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", returnDate)
	router.Run("localhost:8080")
}

func returnDate(c *gin.Context) {
	t := time.Now()
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Hello Gin",
		"year":  t.Year(),
		"month": int(t.Month()),
		"day":   t.Day(),
	})
}

