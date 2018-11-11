package data

import (
	"../models"
	"../database"
)

func GetAll() []models.GoTest {
	var GoTests []models.GoTest
	var err error
	rows,err := database.DBCon.Query(`
		SELECT
			id,
			date,
			created_on,
			tes
		FROM go_test
		ORDER BY id DESC`)

	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		repo := models.GoTest{}
		rows.Scan(
			&repo.Id,
			&repo.Date,
			&repo.CreatedOn,
			&repo.Tes,
		)

		GoTests = append(GoTests, repo)
	}
	return GoTests
}

func GetAllbyID(id string) (models.GoTest,error) {
    rows := database.DBCon.QueryRow(`
			SELECT
				id,
				date,
				created_on,
				tes
			FROM go_test
			WHERE id=$1
			ORDER BY id DESC`, id)

		repo := models.GoTest{}
    err := rows.Scan(
					&repo.Id,
					&repo.Date,
					&repo.CreatedOn,
					&repo.Tes,
				)

    return repo,err
}
