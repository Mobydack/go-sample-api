package main

import "github.com/gin-gonic/gin"

import "tic-tac-toe-api/api"

func main() {
	router := gin.Default()

	ApiRoutes := router.Group("api")

	{
		ApiRoutes.GET("scores", api.GetScores)
	}

	router.Run(":3000")
}
