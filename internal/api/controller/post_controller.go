package controller

import (
	"mindlab/internal/model"
	"mindlab/internal/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Repo repo.PostRepository
}

func NewPostController(repo repo.PostRepository) *PostController {
	return &PostController{Repo: repo}
}

func (c *PostController) Create(ctx *gin.Context) {
	var u model.Post
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := c.Repo.Create(ctx, &u); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

func (c *PostController) Index(ctx *gin.Context) {
	users, err := c.Repo.Index(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (c *PostController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var p model.Post
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.ID = id

	if _, err := c.Repo.Update(ctx, &p); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func (c *PostController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.Repo.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func (c *PostController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := c.Repo.Show(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, post)
	}
}
