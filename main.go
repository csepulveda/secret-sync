package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/csepulveda/secret-sync/config"
	"github.com/csepulveda/secret-sync/sync"
	"github.com/csepulveda/secret-sync/sync/k8s"
)

func main() {
	configPath := "/etc/config/config.json"

	sleep, exists := os.LookupEnv("INTERVAL")
	if !exists {
		sleep = "120"
	}

	intSleep, err := strconv.Atoi(sleep)
	if err != nil {
		log.Printf("bar interval env_var: %v", err)
		panic("oh no")
	}

	for {

		log.Printf("running every %d seconds\n", intSleep)
		cfg, err := config.ReadConfig(configPath)
		if err != nil {
			log.Printf("error reading config: %v", err)
			panic("oh no")
		}

		//Create secrets
		for i := range cfg.Secrets {
			log.Printf("Sync %d of %d secrets\n", i+1, len(cfg.Secrets))
			err := sync.SyncSecret(cfg.Secrets[i])
			if err != nil {
				log.Printf("error reading config: %v", err)
			}
		}

		//Delete secrets by sync_secret but not defined in the actual configuration
		k8s.DeleteSecrets(cfg)
		time.Sleep(time.Second * time.Duration(intSleep))
	}

}
