package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done!in path " + cCp.Request.URL.Path)
		}()
		c.String(200, "ddd")
	})

	router.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done in path" + c.Request.URL.Path)
	})
	_ = router.Run(":8084")
}
