package controller

import (
	"fil303/restApi/model"
	"github.com/gin-gonic/gin"
)

func GetAllAlbums(c *gin.Context) {
	albums := model.GetAllAlbums()
	c.JSON(200, gin.H{
		"message" : "Get All Deta Successfully",
		"data" : albums,
	})
}

func GetAlbumById(c *gin.Context){
    albam  := model.GetAlbumById(c.Param("id"))
	c.JSON(200,gin.H{
		"message" : "Album Get Successfly",
		"album" : albam,
	})
}

func AddAlbum(c *gin.Context){
	var album model.Album
	err := c.BindJSON(&album)
	if err == nil {}
	c.JSON(200,gin.H{
		"message" : "Album Get Successfly",
		"data" : model.AddAlbum(album),
	})
}