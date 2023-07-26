package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/sidgim/usergo/awsgo"
	"github.com/sidgim/usergo/models"
)

// Get Keys of postgres with secret Manager
func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}
	json.Unmarshal(key.SecretBinary, &secretData)
	fmt.Println("Reading ok " + nameSecret)
	return secretData, nil
}
