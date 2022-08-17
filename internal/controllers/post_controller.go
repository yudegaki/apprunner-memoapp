package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/external"
	"github.com/yudegaki/apprunner-memoapp/internal/repositories"
	"github.com/yudegaki/apprunner-memoapp/internal/usecases"
)

type PostResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func convertEntityPostToControllerPost(entityPost *entities.Post) *PostResponse {
	return &PostResponse{
		Title: entityPost.GetTitle(),
		Body:  entityPost.GetBody(),
	}
}

func convertEntityPostsToControllerPosts(entityPosts []*entities.Post) []*PostResponse {
	var controllerPosts []*PostResponse
	for _, post := range entityPosts {
		controllerPosts = append(controllerPosts, convertEntityPostToControllerPost(post))
	}
	return controllerPosts
}

func GetAllPosts(c *gin.Context) {
	repository := repositories.NewPostRepository(external.DB)
	usecase := usecases.NewGetPostsUsecase(repository)

	result, err := usecase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := convertEntityPostsToControllerPosts(result)
	c.JSON(200, resp)
}

func GetPostDetail(c *gin.Context) {
	repository := repositories.NewPostRepository(external.DB)
	usecase := usecases.NewGetPostUsecase(repository)

	var id uint
	if _id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		id = uint(_id)
	}

	result, err := usecase.Execute(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := convertEntityPostToControllerPost(result)
	c.JSON(200, resp)
}

func CreatePost(c *gin.Context) {
	repository := repositories.NewPostRepository(external.DB)
	usecase := usecases.NewCreatePostUsecase(repository)

	var req PostResponse
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var post entities.Post
	post.SetTitle(req.Title)
	post.SetBody(req.Body)

	if err := usecase.Execute(post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func UpdatePost(c *gin.Context) {
	repository := repositories.NewPostRepository(external.DB)
	usecase := usecases.NewUpdatePostUsecase(repository)

	var req PostResponse
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var post entities.Post
	post.SetTitle(req.Title)
	post.SetBody(req.Body)

	if err := usecase.Execute(post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func DeletePost(c *gin.Context) {
	repository := repositories.NewPostRepository(external.DB)
	usecase := usecases.NewDeletePostUsecase(repository)

	var id uint
	if _id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		id = uint(_id)
	}

	if err := usecase.Execute(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
