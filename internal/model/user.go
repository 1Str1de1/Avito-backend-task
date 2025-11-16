package model

import "errors"

var (
	ErrEmptyUserName  = errors.New("username can't be empty")
	ErrEmptyTeamName  = errors.New("team name can't be empty")
	ErrNoRowsAffected = errors.New("rows were not changed successfully")
)

type User struct {
	UserId   string `json:"id" db:"userId"`
	Username string `json:"username" db:"username" binding:"required"`
	TeamName string `json:"team_name" db:"team_name"`
	IsActive bool   `json:"is_active" db:"is_active" binding:"required"`
}

func (db *DB) CreateUser(u *User) error {
	if len(u.Username) == 0 {
		return ErrEmptyUserName
	}
	if len(u.TeamName) == 0 {
		return ErrEmptyTeamName
	}
	sqlStatement := `
	INSERT INTO users (userid, username, team_name, is_active)
	VALUES ($1, $2, $3, $4);`

	res, err := db.db.Exec(sqlStatement, u.UserId, u.Username, u.TeamName, u.IsActive)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return ErrNoRowsAffected
	}

	return nil
}

func (db *DB) SetIsActive(u *User) error {
	sqlStatement := `
	UPDATE users
	SET is_active=$1
	WHERE userid=$2;`

	res, err := db.db.Exec(sqlStatement, u.IsActive, u.UserId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return ErrNoRowsAffected
	}

	return nil
}

func (db *DB) GetUserById(id int) (*User, error) {
	sqlStatement := `
	SELECT userid, username, team_name, is_active
	FROM users
	WHERE userid=$1;`

	var u User
	row := db.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&u.UserId, &u.Username, &u.TeamName, &u.IsActive)

	if err != nil {
		return nil, err
	}

	return &u, nil

}
