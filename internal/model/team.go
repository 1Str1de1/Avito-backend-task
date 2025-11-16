package model

type Team struct {
	TeamName string `json:"team_name" db:"team_name" binding:"required"`
	Members  []User `json:"members" binding:"required,min=1"`
}

func (db *DB) CreateTeam(t *Team) error {
	if len(t.TeamName) == 0 {
		return ErrEmptyTeamName
	}
	sqlStatement := `
	INSERT INTO team (team_name)
	VALUES ($1);`

	res, err := db.db.Exec(sqlStatement, t.TeamName)
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

func (db *DB) GetTeam(teamName string) (*Team, error) {
	sqlStatement := `SELECT * FROM team WHERE team_name=$1;`

	row := db.db.QueryRow(sqlStatement, teamName)

	err := row.Scan(&teamName)
	if err != nil {
		return nil, err
	}

	team := &Team{TeamName: teamName}

	err = row.Err()
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (db *DB) GetTeamUsers(teamName string) ([]User, error) {
	sqlStatement := `SELECT * FROM users WHERE team_name=$1`

	users := make([]User, 0, 20)

	rows, err := db.db.Query(sqlStatement, teamName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, username string
		var isActive bool
		err = rows.Scan(&id, &username, &teamName, &isActive)
		if err != nil {
			// TODO: error handling
			return nil, err
		}
		users = append(users, User{
			UserId:   id,
			Username: username,
			TeamName: teamName,
			IsActive: isActive,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (db *DB) GetAllTeams() ([]*Team, error) {
	sqlStatement := `SELECT * FROM team LIMIT $1;`
	teams := make([]*Team, 0, 20)
	rows, err := db.db.Query(sqlStatement, 20)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var teamName string
		err = rows.Scan(&teamName)
		if err != nil {
			return nil, err
		}

		team := &Team{TeamName: teamName}
		teams = append(teams, team)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return teams, nil
}
