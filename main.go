package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func main() {
	logger := logrus.New()

	logFile, err := os.OpenFile("/logs/go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.WithError(err).Info("Failed to log to file, using default stderr")
	}
	logger.SetOutput(logFile)

	logger.SetFormatter(&ecslogrus.Formatter{})

	logger.SetReportCaller(true)
	logger.SetLevel(logrus.TraceLevel)

	logger.Infoln("Hello World")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		logger.Infoln("ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
