package handler

import (
	"fmt"
	"github.com/1Str1de1/Avito-backend-task/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HandlePrCreate(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			PrName   string `json:"pr_name" binding:"required"`
			AuthorId int    `json:"author_id" binding:"required"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("bad request: %v", err),
			})
			return
		}

		if body.AuthorId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("invalid user id:"),
			})
			return
		}

		if len(body.PrName) == 0 {
			c.JSON(http.StatusBadRequest, model.ErrorEmptyPrName())
			return
		}

		reviewers, err := db.AssignAuthors(body.AuthorId)

		prId := db.GeneratePrIdWithUuid()
		pr := model.PullRequest{
			PrId:              prId,
			PrName:            body.PrName,
			AuthorId:          body.AuthorId,
			PrStatus:          model.OPEN,
			AssignedReviewers: reviewers,

			CreatedAt: time.Now(),
		}

		err = db.CreatePR(&pr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":            http.StatusCreated,
			"pull_request_id":   pr.PrId,
			"pull_request_name": body.PrName,
			"author_id":         body.AuthorId,
			"message":           "PR created successfully",
		})
	}
}

func HandleMergeRequest(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func HandleReassignAuthor(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}
