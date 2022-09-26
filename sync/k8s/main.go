package k8s

import (
	"context"
	"encoding/json"
	"log"

	"github.com/csepulveda/secret-sync/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func CreateSecret(namespace, secretname, secret string) {
	kind := "Secret"
	apiVersion := "v1"

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	secretData := map[string]string{}
	err = json.Unmarshal([]byte(secret), &secretData)
	if err != nil {
		panic(err.Error())
	}

	labels := map[string]string{}
	labels["created_by"] = "secret-sync"

	k8ssecret := &corev1.SecretApplyConfiguration{
		TypeMetaApplyConfiguration: v1.TypeMetaApplyConfiguration{
			Kind:       &kind,
			APIVersion: &apiVersion,
		},
		ObjectMetaApplyConfiguration: &v1.ObjectMetaApplyConfiguration{
			Name:      &secretname,
			Namespace: &namespace,
			Labels:    labels,
		},
		StringData: secretData,
	}
	opts := metav1.ApplyOptions{
		FieldManager: "secret-sync",
	}
	_, err = clientset.CoreV1().Secrets(namespace).Apply(context.TODO(), k8ssecret, opts)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteSecrets(cfg *config.Config) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	wantedSecrets := Secrets{}
	actualSecrets := Secrets{}

	for i := range cfg.Secrets {
		wantedSecret := Secret{
			Name:      cfg.Secrets[i].Dest,
			Namespace: cfg.Secrets[i].Namespace,
		}
		wantedSecrets.AddSecret(wantedSecret)

		namespace := cfg.Secrets[i].Namespace
		opts := metav1.ListOptions{
			LabelSelector: "created_by=secret-sync",
		}

		secretList, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), opts)
		if err != nil {
			panic(err.Error())
		}

		for _, secret := range secretList.Items {
			actualSecret := Secret{
				Name:      secret.Name,
				Namespace: secret.Namespace,
			}
			actualSecrets.AddSecret(actualSecret)
		}
	}

	for _, actual := range actualSecrets.Secrets {
		toDelete := true
		for _, wanted := range wantedSecrets.Secrets {
			if actual == wanted {
				toDelete = false
			}
		}
		if toDelete {
			err := clientset.CoreV1().Secrets(actual.Namespace).Delete(context.TODO(), actual.Name, metav1.DeleteOptions{})
			if err != nil {
				log.Printf("error deleting secret %s on namespace %s. error: %v", actual.Name, actual.Namespace, err)
			} else {
				log.Printf("deleted secret %s on namespace %s. error: %v", actual.Name, actual.Namespace, err)
			}
		}
	}

}
