package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rsaldano/rds-killmysql/awsgo"
	"github.com/rsaldano/rds-killmysql/models"
)

//GetSecret es la función que devuelve la password de Secret Manager
func GetSecret(secretName string, aws_region string) (datos models.SecretRDSJson) {
	var datosSecret models.SecretRDSJson

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		panic("Error al leer el valor de Secret Manager " + err.Error())
	}
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("Lectura correcta de secret :")
	return datosSecret
}
