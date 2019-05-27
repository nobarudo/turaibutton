package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.Static("/public", "./public")
	router.LoadHTMLGlob("views/*")

	router.GET("/", indexGET)
	router.GET("/turai", turaiGET)
	router.Run(":8080")
}

func indexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func turaiGET(c *gin.Context) {
	turaiText := "つらい"
	lines := []string{}
	f, err := os.Open("./turai.txt")
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	length := len(lines)
	num := rand.Intn(length)

	log.Println(lines[num])
	turaiText = lines[num]

	c.String(http.StatusOK, turaiText)
}
