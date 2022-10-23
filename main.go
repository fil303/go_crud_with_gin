package main

import (
	"github.com/gin-gonic/gin"
	"fil303/restApi/controller"
)


func main() {
	r := gin.Default()
	r.GET("/albums", controller.GetAllAlbums)
	r.POST("/albums", controller.AddAlbum)
	r.GET("/albums/:id", controller.GetAlbumById)
	r.Run("localhost:8080")
}
