package controllers

import (
	"github.com/Itsu8/rest_api/initializers"
	"github.com/Itsu8/rest_api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Title}
	result := initializers.DB.Create(&post)
	// Return it
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsRead(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id off URL
	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}
