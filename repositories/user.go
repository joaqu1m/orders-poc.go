package repositories

import (
	"orders-api/database"
	"orders-api/dtos"
	"orders-api/models"
)

func GetUsers() ([]string, error) {
	rows, err := database.MySQLQuery("SELECT login FROM `user`")
	if err != nil {
		return nil, err
	}

	var users []string
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Login); err != nil {
			return nil, err
		}
		users = append(users, user.Login)
	}

	return users, nil
}

func CreateUser(user dtos.UserRequest) error {
	_, err := database.MySQLQuery("INSERT INTO `user` (login, password) VALUES (?, ?)", user.Login, user.Password)
	if err != nil {
		return err
	}

	_, err = database.MongoInsertOne("users", user)
	if err != nil {
		return err
	}

	return nil
}

func AuthUser(user dtos.UserRequest) (bool, error) {
	row, err := database.MySQLQueryRow("SELECT login FROM `user` WHERE login = ? AND password = ?", user.Login, user.Password)
	if err != nil {
		return false, err
	}

	if err := row.Scan(&user.Login); err != nil {
		return false, nil
	}

	return true, nil
}
