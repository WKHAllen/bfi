package main

import (
	"net/http"
	"os"
	"math/rand"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

const interpretTimeout = 10 * time.Second
var sessions map[int]*BFInterpreter

func index(c *gin.Context) {
	newSessionID := rand.Int()
	for _, exists := sessions[newSessionID]; exists; {
		newSessionID = rand.Int()
	}
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"sessionID": newSessionID,
	})
}

func doInterpret(bfi *BFInterpreter) (bool, int, byte, error) {
	var returnCode int
	var displayByte byte
	var err error

	done := make(chan int)
	go func() {
		returnCode, displayByte, err = bfi.Interpret()
		done <- 0
	}()

	select {
		case <-done:
			return true, returnCode, displayByte, err
		case <-time.After(interpretTimeout):
			return false, returnCode, displayByte, err
	}
}

func interpret(c *gin.Context) {
	sessionID, err := strconv.Atoi(c.Query("sessionID"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid sessionID",
		})
	} else {
		code := c.Query("code")
		bfi := NewBFInterpreter(code)
		sessions[sessionID] = bfi

		success, returnCode, displayByte, err := doInterpret(bfi)

		if !success {
			delete(sessions, sessionID)
			c.JSON(http.StatusOK, gin.H{
				"error": "timed out",
			})
		} else {
			if err != nil {
				delete(sessions, sessionID)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
					"index": bfi.index,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"returnCode": returnCode,
					"displayByte": displayByte,
				})
			}
		}
	}
}

func returnOutput(c *gin.Context) {
	// sessionID := c.Query("sessionID")
	// TODO: continue interpreting
}

func returnInput(c *gin.Context) {
	// sessionID := c.Query("sessionID")
	// value := c.Query("value")
	// TODO: pass the input value to the interpreter
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	rand.Seed(time.Now().UTC().UnixNano())
	
	sessions = make(map[int]*BFInterpreter)

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", index)
	router.GET("/interpret", interpret)
	router.GET("/returnInput", returnInput)

	router.Run(":" + port)
}
