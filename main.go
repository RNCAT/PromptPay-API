package main

import (
	"github.com/RNCAT/PromptPay/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/promptpay", handlers.CreatePromptPay)
	r.Run(":3001")
}
