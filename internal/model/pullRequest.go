package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

const (
	OPEN   PrStatus = "OPEN"
	MERGED PrStatus = "MERGED	"
)

type PullRequest struct {
	PrId              string
	PrName            string
	AuthorId          int
	PrStatus          PrStatus
	AssignedReviewers []int // UserIds of assigned reviewers
	NeedMoreReviewers bool
	CreatedAt         time.Time
	MergedAt          time.Time
}

type PullRequestShort struct {
	PrId     string
	PrName   string
	AuthorId string
	PrStatus PrStatus
}

type PrStatus string

func (db *DB) GeneratePrIdWithUuid() string {
	id := uuid.New().String()[:8]
	return fmt.Sprintf("pr-%s", id)
}

func (db *DB) CreatePR(pr *PullRequest) error {
	if len(pr.PrName) == 0 {
		return errors.New("PR name can't be empty") // TODO: autonaming
	}
	sqlStatement := `
	INSERT INTO pull_request (
		name, author_id, reviewers_id 
	)
	VALUES($1, $2, $3);`

	res, err := db.db.Exec(sqlStatement, pr.PrName,
		pr.AuthorId, pq.Array(pr.AssignedReviewers))

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return err
	}

	return nil

}

func (db *DB) AssignAuthors(authorId int) ([]int, error) {
	sqlStatement := `
	SELECT userid
	FROM users
	WHERE is_active=true
	  AND userid!=$1
	  AND team_name = (
    	SELECT team_name 
    	FROM users 
   		WHERE userid = $1
  )
	LIMIT 2;`

	rows, err := db.db.Query(sqlStatement, authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviewers := make([]int, 0, 2)

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			// TODO: log error
			continue
		}

		reviewers = append(reviewers, id)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return reviewers, nil
}
