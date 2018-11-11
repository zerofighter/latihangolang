package data

import (
	"../models"
	"../database"
	"fmt"
)

func GetAll() []models.GoTest {
	var GoTests []models.GoTest
	var err error
	rows,err := database.DBCon.Query(`
		SELECT
			id,
			first_name,
			last_name,
			place_of_birth,
			date_of_birth,
			marital_status,
			blood_type,
			gender,
			religion,
			phone,
			mobile_phone,
			email,
			citizen_id_address,
			postal_code,
			residential_address,
			same_with_citizen_id_address,
			identity_type,
			identity_number,
			expired_date_identity,
			no_expired_date_identity
		FROM applicants_personal_data
		ORDER BY id DESC`)

	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		repo := models.GoTest{}
		rows.Scan(
			&repo.Id,
			&repo.First_name,
			&repo.Last_name,
			&repo.Place_of_birth,
			&repo.Date_of_birth,
			&repo.Marital_status,
			&repo.Blood_type,
			&repo.Gender,
			&repo.Religion,
			&repo.Phone,
			&repo.Mobile_phone,
			&repo.Email,
			&repo.Citizen_id_address,
			&repo.Postal_code,
			&repo.Residential_address,
			&repo.Same_with_citizen_id_address,
			&repo.Identity_type,
			&repo.Identity_number,
			&repo.Expired_date_identity,
			&repo.No_expired_date_identity,
		)
		fmt.Println(repo.First_name)
		GoTests = append(GoTests, repo)
	}
	return GoTests
}

func GetAllbyID(id string) (models.GoTest,error) {
    rows := database.DBCon.QueryRow(`
			SELECT
					id,
					first_name,
					last_name,
					place_of_birth,
					date_of_birth,
					marital_status,
					blood_type,
					gender,
					religion,
					phone,
					mobile_phone,
					email,
					citizen_id_address,
					postal_code,
					residential_address,
					same_with_citizen_id_address,
					identity_type,
					identity_number,
					expired_date_identity,
					no_expired_date_identity
			FROM applicants_personal_data
			WHERE id=$1
			ORDER BY id DESC`, id)

		repo := models.GoTest{}
    err := rows.Scan(
					&repo.Id,
					&repo.First_name,
					&repo.Last_name,
					&repo.Place_of_birth,
					&repo.Date_of_birth,
					&repo.Marital_status,
					&repo.Blood_type,
					&repo.Gender,
					&repo.Religion,
					&repo.Phone,
					&repo.Mobile_phone,
					&repo.Email,
					&repo.Citizen_id_address,
					&repo.Postal_code,
					&repo.Residential_address,
					&repo.Same_with_citizen_id_address,
					&repo.Identity_type,
					&repo.Identity_number,
					&repo.Expired_date_identity,
					&repo.No_expired_date_identity,
				)

    return repo,err
}

func Update(data models.GoTest) error {
		fmt.Println(data.First_name)
    return nil
}
