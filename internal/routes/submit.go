package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubmitRequestBody struct {
	Code string `json:"code"`
}

func SubmitRoutes(routerGroup *gin.RouterGroup) {
	r := routerGroup.Group("/api")
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// Endpoint to get the uploaded code file
		r.POST("/submit", func(c *gin.Context) {
			// single file
			log.Println(c.ContentType())
			file, err := c.FormFile("file")
			if err != nil {
				log.Printf("%v", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "'file=@/path/to/your/file.txt'; the parameter must be 'file'",
				})
				return
			}

			// Stream the file to the worker via gRPC somehow
			// https://dev.to/dialaeke/streaming-large-files-between-microservices-a-grpc-implementation-5ee8

			c.JSON(http.StatusOK, gin.H{
				"pls_work": file.Filename,
			})
		})
	}
}
