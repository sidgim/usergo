package bd

import (
	"database/sql"
	"fmt"
	"github.com/sidgim/usergo/models"
	"github.com/sidgim/usergo/secretm"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("postgres", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexión exitosa de la BD")
	return nil
}

func ConnStr(key models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = key.Username
	authToken = key.Password
	dbEndpoint = key.Host
	dbName = "waypoint"
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", dbEndpoint, dbUser, authToken, dbName)
}
