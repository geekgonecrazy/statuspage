package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/eirsyl/statuspage/statuspage"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"log"
)

func main()  {

	redisAddr := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")

	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		log.Print("Could not parse Redis DB, using 0 as default.")
		redisDBInt = 0
	}


	db := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDBInt,
	})

	services := statuspage.Services{}
	services.Initialize(*db)

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"owner": "Abakus",
			"services": services.GetServices(),
		})
	})

	router.Run()

}