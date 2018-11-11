package models

type (
	GoTest struct {
		Id        										int 		`form:"id" json:"id"`
		First_name      							string	`form:"first_name" json:"first_name"`
		Last_name      								string 	`form:"last_name" json:"last_name"`
		Place_of_birth      					string	`form:"place_of_birth" json:"place_of_birth"`
		Date_of_birth									string	`form:"date_of_birth" json:"date_of_birth"`
		Marital_status								int			`form:"mrital_status" json:"marital_status"`
		Blood_type										int			`form:"blood_type" json:"blood_type"`
		Gender												int			`form:"gender" json:"gender"`
		Religion											int			`form:"religion" json:"religion"`
		Phone													string	`form:"phone" json:"phone"`
		Mobile_phone									string	`form:"mobile_phone" json:"mobile_phone"`
		Email													string	`form:"email" json:"email"`
		Citizen_id_address						string	`form:"citizen_id_address" json:"citizen_id_address"`
		Postal_code										int			`form:"postal_code" json:"postal_code"`
		Residential_address						string	`form:"residential_address" json:"residential_address"`
		Same_with_citizen_id_address	int			`form:"same_with_citizen_id_address" json:"same_with_citizen_id_address"`
		Identity_type									int			`form:"identity_type" json:"identity_type"`
		Identity_number								string	`form:"identity_number" json:"identity_number"`
		Expired_date_identity					string	`form:"expired_date_identity" json:"expired_date_identity"`
		No_expired_date_identity			int			`form:"no_expired_date_identity" json:"no_expired_date_identity"`
	}
)
