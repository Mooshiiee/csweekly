package queries

import (
	"database/sql"
	"fmt"
	"os"

	"csweeklyproj/models"
)

// takes in sql.db object, will return an array of users along with error
func QueryUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query : %v\n", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
			//return users as nil, and error as the formatted message

		}

		users = append(users, user)
		fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
		//return users as nil, and error as the formatted message
	}

	return users, nil
	//return users and nil for the error.
}

func QueryProblems(db *sql.DB) ([]models.Problem, error) {
	rows, err := db.Query("SELECT * FROM problems")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//initiate a problem slice to hold data from the rows
	var problems []models.Problem

	//loop through rows and use Scan() to assign db data to struct data
	for rows.Next() {
		//variable to hold instance data
		var problem models.Problem
		if err := rows.Scan(&problem.ID, &problem.Title,
			&problem.Text, &problem.Constraints, &problem.Hint,
			&problem.Solution); err != nil {
			return problems, err
		}
		problems = append(problems, problem)
	}
	if err = rows.Err(); err != nil {
		return problems, err
	}
	return problems, err
}

// takes in an sql.DB connection and the id from the request url, then returns
func QuerySingleProblem(db *sql.DB, id int) (models.Problem, error) {

	var problem models.Problem

	err := db.QueryRow("SELECT * FROM problems WHERE id = ?", id).Scan(
		&problem.ID,
		&problem.Title,
		&problem.Text,
		&problem.Hint,
		&problem.Constraints,
		&problem.Solution,
		&problem.IsProject,
	)
	if err != nil {
		//return an empty struct if error
		return models.Problem{}, fmt.Errorf("error scanning Problem instance: %w", err)
	}

	return problem, nil
}
