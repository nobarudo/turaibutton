package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	access_log, err := os.OpenFile("access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}
	gin.DefaultWriter = io.MultiWriter(access_log)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
