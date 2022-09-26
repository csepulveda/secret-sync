package sync

import (
	"fmt"

	"github.com/csepulveda/secret-sync/config"
	"github.com/csepulveda/secret-sync/sync/aws"
)

func SyncSecret(secret *config.Secret) error {
	switch provider := secret.Provider; provider {
	case "aws":
		_, err := aws.SyncSecret(secret)
		return err
	default:
		err := fmt.Errorf("provider %q not supported", provider)
		return err
	}
}
