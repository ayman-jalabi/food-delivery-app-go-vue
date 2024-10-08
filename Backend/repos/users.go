package repos

import (
	"database/sql"
	"fmt"
	"log"
	"main/models"
)

type UserRepo struct {
	//you can add here the db open function. you can find this example on the documentation maybe page:
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{db}
}

func (ur UserRepo) CheckEmailExistence(email string) (bool, error) {
	emailExists := false
	var emailResult string

	// Execute the query to search for the email in the database
	result, err := ur.Db.Query("SELECT email FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println("error checking if email exists or not: ", err)
		return false, err
	}
	defer result.Close()

	// Iterate over the results (there should only be one, but still a good practice)
	for result.Next() {
		// Scan the result into the emailResult variable
		err = result.Scan(&emailResult)
		if err != nil {
			fmt.Println("error scanning result: ", err)
			return false, err
		}

		// Check if the returned email matches the input email
		if emailResult == email {
			emailExists = true
			break
		}
	}

	// Check for errors encountered during iteration
	err = result.Err()
	if err != nil {
		fmt.Println("error during result iteration: ", err)
		return false, err
	}

	return emailExists, nil
}

func (ur UserRepo) CreateUser(user models.User) error {
	_, err := ur.Db.Exec(
		"INSERT INTO users (email, first_name, last_name, address, phone_number, password) values ($1, $2, $3, $4, $5, $6)",
		user.Email, user.FirstName, user.LastName, user.Address, user.PhoneNumber, user.Password,
	)

	return err
}

func (ur UserRepo) GetUser(email string) (models.User, error) {

	var user models.User

	result, err := ur.Db.Query("SELECT email, first_name, last_name, address, phone_number FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println("error getting user: ", err)
		return user, err
	}
	for result.Next() {
		user = models.User{}
		err = result.Scan(&user.ID)
		if err != nil {
			fmt.Println("error scanning result: ", err)
			return user, err
		}
	}
	result.Close()

	return user, nil
}

func (ur UserRepo) DeleteUser(id int) error {
	stmt, err := ur.Db.Prepare("DELETE FROM users WHERE id = $1;")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	log.Printf("ID = %d, number of affected rows= %d\n", id, rowCnt)

	return err
}
