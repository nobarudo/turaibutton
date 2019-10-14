package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("access.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

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

	turaiText = lines[num]

	c.String(http.StatusOK, turaiText)
}
