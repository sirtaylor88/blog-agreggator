package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	dirpath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error finding HOME dir: %w", err)
	}
	return fmt.Sprintf("%s/%s", dirpath, configFileName), nil
}

func write(config Config) error {
	fp, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("Error marshaling config data: %w", err)
	}

	if err = os.WriteFile(fp, data, 0644); err != nil {
		return fmt.Errorf("Error writing to config file: %w", err)
	}
	return nil
}

func Read() (Config, error) {
	var config Config

	fp, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(fp)

	if err != nil {
		return Config{}, fmt.Errorf("Error reading config file: %w", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Error unmarshaling json data: %w", err)
	}

	return config, nil
}

func SetUser(username string, config Config) error {
	config.CurrentUserName = username

	if err := write(config); err != nil {
		return fmt.Errorf("Error setting current user to config file: %w", err)
	}

	return nil
}
