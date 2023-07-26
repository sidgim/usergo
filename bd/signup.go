package bd

import (
	"fmt"
	"github.com/sidgim/usergo/models"
	"github.com/sidgim/usergo/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	r := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) Values ( " +
		"'" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DatePostgreSQL() + "')"
	fmt.Println(r)
	_, err = Db.Exec(r)
	if err != nil {
		return err
	}
	return nil
}
