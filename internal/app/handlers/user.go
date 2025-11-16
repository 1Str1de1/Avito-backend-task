package handler

import (
	"fmt"
	"github.com/1Str1de1/Avito-backend-task/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleUpdateIsActive(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			UserID   int  `json:"userid" binding:"required"`
			IsActive bool `json:"is_active" binding:"required"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("bad request: %v", err),
			})
			return
		}

		user, err := db.GetUserById(body.UserID)
		if err != nil {
			c.JSON(http.StatusNotFound, model.ErrorNotFound())
			return
		}

		user.IsActive = body.IsActive
		err = db.SetIsActive(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": fmt.Sprintf("failed to update user status: %v", err),
			})
			return
		}

		updatedUser, err := db.GetUserById(body.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": fmt.Sprintf("failed to fetch updated user: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "user status updated successfully",
			"user":    updatedUser,
		})
	}

}

func HandleGetUser(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Param("id")
		id, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("invalid user id: %v", err),
			})
			return
		}

		user, err := db.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, model.ErrorNotFound())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   user,
		})
	}
}

func HandleGetUserPrs(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}
