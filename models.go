package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbDriver   = "mysql"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "mysql"
	dbUser     = "root"
	dbPassword = "secret"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func newUser(firstName, lastName, email, password string) User {
	return User{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashAndSalt(password),
	}
}

func (user User) save() error {
	ds := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, ds)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println("Successfully connected to db")

	query := fmt.Sprintf("INSERT INTO users (id, first_name, last_name, email, password) VALUES ('%v', '%v', '%v', '%v', '%v')", user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	data, err := db.Query(query)
	if err != nil {
		return err
	}
	defer data.Close()
	fmt.Println("Successfully created User in db")

	return nil
}
