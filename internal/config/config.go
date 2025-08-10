package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}


func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	if err := os.WriteFile(configPath, jsonData, 0644); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}
	return nil
}

func ReadConfig() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file content: %w", err)
	}
	
	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return Config{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, configFileName)

	return configPath, nil
}
