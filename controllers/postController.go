package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/repository/initializers"
	"github.com/repository/models"
)

func PostCreate(c *gin.Context) {
	/* Get data of request body */
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	/* Save the new post  */
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	/* Return response */
	c.JSON(200, gin.H{
		"status": "ok",
		"post":   post,
	})
}

func PostGetAll(c *gin.Context) {
	/* Search in DB */
	var posts []models.Post
	initializers.DB.Find(&posts)

	/* Return response */
	c.JSON(200, gin.H{
		"status": "ok",
		"posts":  posts,
	})
}

func PostGet(c *gin.Context) {
	/* Get the ID from Params URL */
	id := c.Param("id")

	/* Search in DB */
	var post models.Post
	initializers.DB.First(&post, id)

	/* Return response */
	c.JSON(200, gin.H{
		"status": "ok",
		"post":   post,
	})
}

func PostUpdate(c *gin.Context) {
	/* Get the ID from Params URL */
	id := c.Param("id")

	/* Get data of request body */
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	/* Find the post were updating */
	var post models.Post
	initializers.DB.First(&post, id)

	/* Update it */
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	/* Return response */
	c.JSON(200, gin.H{
		"status": "ok",
		"post":   post,
	})
}

func PostDelete(c *gin.Context) {
	/* Get the ID from Params URL */
	id := c.Param("id")

	/* Delete the post */
	initializers.DB.Delete(&models.Post{}, id)

	/* Return response */
	c.Status(200)
}
