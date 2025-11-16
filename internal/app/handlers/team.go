package handler

import (
	"fmt"
	"github.com/1Str1de1/Avito-backend-task/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleAddTeam(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			TeamName string       `json:"team_name" binding:"required"`
			Members  []model.User `json:"members" binding:"required,min=1"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("bad request: %v", err),
			})
			return
		}

		if len(body.TeamName) == 0 {
			c.JSON(http.StatusBadRequest, model.ErrorEmptyTeamName())
			return
		}

		if len(body.Members) == 0 {
			c.JSON(http.StatusBadRequest, model.NoUsersAdded)
			return
		}

		team := model.Team{
			TeamName: body.TeamName,
		}

		err := db.CreateTeam(&team)

		if err != nil {
			c.JSON(http.StatusBadRequest, model.ErrorTeamExists())
			return
		}

		membersAdded := 0

		for _, member := range body.Members {
			user := &model.User{
				UserId:   member.UserId,
				Username: member.Username,
				TeamName: body.TeamName,
				IsActive: member.IsActive,
			}

			err := db.CreateUser(user)
			if err != nil {
				// TODO: log error
				continue
			}
			membersAdded++
		}

		if membersAdded == 0 {
			c.JSON(http.StatusInternalServerError, model.ErrorNoUsersAdded())
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":    http.StatusCreated,
			"team_name": body.TeamName,
			"members":   body.Members,
			"message":   "team created successfully with members",
		})
	}
}

func HandleGetTeam(db *model.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		teamName := c.Query("team_name")

		if len(teamName) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "team name query parameter is required",
			})
			return
		}

		team, err := db.GetTeam(teamName)
		if team == nil || err != nil {
			c.JSON(http.StatusNotFound, model.ErrorNotFound())
			return
		}

		members, err := db.GetTeamUsers(team.TeamName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  fmt.Sprintf("error fetching members %v", err),
			})
			return
		}

		team.Members = members

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   team,
		})
	}
}
