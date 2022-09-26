package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadConfig(path string) (*Config, error) {

	safePath := filepath.Clean(path)
	if !strings.HasPrefix(safePath, "/etc/config/") {
		panic(fmt.Errorf("unsafe input"))
	}
	jsonFile, err := os.Open(safePath)
	if err != nil {
		return nil, err
	}
	byteValue, _ := io.ReadAll(jsonFile)
	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Printf("Error unmarshal config file: %v\n", err)
		return nil, err
	}
	return &config, nil
}
