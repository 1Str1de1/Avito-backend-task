package model

import (
	"net/http"
	"strconv"
)

type ErrorCode string

const (
	EmptyTeamName ErrorCode = "EMPTY_TEAM_NAME"
	NoUsersAdded  ErrorCode = "NO_USERS_ADDED"
	TeamExists    ErrorCode = "TEAM_EXISTS"
	EmptyPrName   ErrorCode = "EMPTY_PR_NAME"
	PrExists      ErrorCode = "PR_EXISTS"
	PrMerged      ErrorCode = "PR_MERGED"
	NotAssigned   ErrorCode = "NOT_ASSIGNED"
	NoCandidate   ErrorCode = "NO_CANDIDATE"
	NotFound      ErrorCode = "NOT_FOUND"
)

// ErrorDetail contains error details
type ErrorDetail struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// ErrorResponse standard error response model
type ErrorResponse struct {
	Status string      `json:"status"`
	Error  ErrorDetail `json:"error"`
}

// NewErrorResponse creates new error response with code and message
func NewErrorResponse(code ErrorCode, status int, message string) ErrorResponse {
	return ErrorResponse{
		Status: strconv.Itoa(status),
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

func ErrorTeamExists() ErrorResponse {
	return NewErrorResponse(TeamExists, http.StatusBadRequest, "team with this name already exists")
}

func ErrorEmptyTeamName() ErrorResponse {
	return NewErrorResponse(EmptyTeamName, http.StatusBadRequest, "team name can't be empty")
}

func ErrorNoUsersAdded() ErrorResponse {
	return NewErrorResponse(NoUsersAdded, http.StatusBadRequest, "no members were added to the team")
}

func ErrorNotFound() ErrorResponse {
	return NewErrorResponse(NotFound, http.StatusNotFound, "resource not found")
}

func ErrorEmptyPrName() ErrorResponse {
	return NewErrorResponse(EmptyPrName, http.StatusBadRequest, "PR name can't be empty")
}

func ErrorPrExists() ErrorResponse {
	return NewErrorResponse(PrExists, http.StatusConflict, "PR id already exists")
}

//func ErrorPrMerged() ErrorResponse {
//	return NewErrorResponse(PrMerged, "Pull request is already merged")
//}

func ErrorNotAssigned() ErrorResponse {
	return NewErrorResponse(NotAssigned, http.StatusBadRequest, "user is not assigned to this pull request")
}

//func ErrorNoCandidate() ErrorResponse {
//	return NewErrorResponse(NoCandidate, "No candidate found")
//}
