package frontend

import "github.com/gin-gonic/gin"

type FrontEnd struct {
}

func (obj FrontEnd) StartListening() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
