package main

import (
	"log"
	"time"

	"github.com/dazai404/elk-logger/pkg/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	l, err := logger.NewFileLogger("info.log")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.Use(ginzap.Ginzap(l.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(l.Logger, true))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
