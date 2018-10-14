package main

import (
    "net/http"

    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    )

func main(){
    router := gin.Default()

    router.Use(static.Serve("/", static.LocalFile("./views", true)))

    api := router.Group("/api")
    {
        api.GET("/", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H {
                "message": "pong",
            })
        })
    }

    api.GET("/jokes", JokeHandler)
    api.POST("/jokes/like/:jokeID", LikeJoke)

    router.Run(":3000")
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"Jokes handler not implemented yet",
  })
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"LikeJoke handler not implemented yet",
  })
}
