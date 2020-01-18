package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

var sessions map[string]BFInterpreter

func index(c *gin.Context) {
	// TODO: initialize a session
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func interpret(c *gin.Context) {
	sessionID := c.Query("sessionID")
	bfcode := c.Query("code")
	// TODO: interpret the brainfuck code
	c.JSON(http.StatusOK, gin.H{})
}

func returnOutput(c *gin.Context) {
	sessionID := c.Query("sessionID")
	// TODO: continue interpreting
}

func returnInput(c *gin.Context) {
	sessionID := c.Query("sessionID")
	value := c.Query("value")
	// TODO: pass the input value to the interpreter
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", index)
	router.GET("/interpret", interpret)
	router.GET("/returnInput", returnInput)

	router.Run(":" + port)
}
