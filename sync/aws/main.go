package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/csepulveda/secret-sync/config"
	"github.com/csepulveda/secret-sync/sync/k8s"
)

func SyncSecret(secret *config.Secret) (bool, error) {
	sess, err := createSession()
	if err != nil {
		log.Printf("Unable create session, %v\n", err)
		return false, err
	}

	svcSecretsManager := secretsmanager.New(sess)
	secretData, err := readSecret(secret.Source, svcSecretsManager)
	if err != nil {
		log.Printf("Unable to get secret data, %v\n", err)
		return false, err
	}

	k8s.CreateSecret(secret.Namespace, secret.Dest, *secretData.SecretString)

	return true, nil
}

func readSecret(secret string, svc secretsmanageriface.SecretsManagerAPI) (*secretsmanager.GetSecretValueOutput, error) {
	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secret),
	}
	secretData, err := svc.GetSecretValue(params)
	if err != nil {
		log.Printf("Unable to get secret data, %v\n", err)
		return secretData, err
	}

	return secretData, nil
}

func createSession() (*session.Session, error) {
	sess, err := session.NewSession()
	if err != nil {
		log.Printf("Unable create session, %v\n", err)
		return nil, err
	}
	return sess, nil
}
