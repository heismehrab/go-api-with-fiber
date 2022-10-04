package repositories

import (
	"log"
	"test-api/config"
	"time"
)

type User struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateUser(user *User) User {
	db := config.ConnectToMysql()

	//  Close connection.
	defer config.DisconnectFromMysql()

	result, err := db.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name,
		time.Now().String()+"@gmail.com",
		time.Now().String(),
	)

	if err != nil {
		log.Fatal(err)
	}

	user.Id, err = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	return *user
}
