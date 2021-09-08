package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		ApiRouter(api)
	}

	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
	// 	Cache:      autocert.DirCache("/var/www/.cache"),
	// }

	r.Run(":7000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// log.Fatal(autotls.RunWithManager(r, &m))
}
