package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 临时版本：只验证工具链可用。第 18 章会替换为正式的分层实现。
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All is well",
		})
	})

	log.Println("server is running on :8080")
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("run server failed: %v", err)
	}
}
