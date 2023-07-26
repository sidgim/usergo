package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sidgim/usergo/awsgo"
	"github.com/sidgim/usergo/bd"
	"github.com/sidgim/usergo/models"
	"os"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	//Cargamos AWS y sus configuraciones
	awsgo.InitAWS()

	if !ValidateParams() {
		fmt.Println("Error in the params. has send 'SecretManager")
		return event, errors.New("error in the params. has send 'SecretManager")
	}
	//Obtenemos data desde Cognito
	var data models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	//Obtenemos
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error reading Secret " + err.Error())
		return event, err
	}
	err = bd.SignUp(data)
	return event, err
}

func ValidateParams() bool {
	_, getParams := os.LookupEnv("SecretName")
	return getParams
}
