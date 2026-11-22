package config

import (
	"fmt"
	"os"
	"strings"
)

type EnvVars struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	dsn        string
	Port       string
}

var envVars *EnvVars

func LoadEnvVars() error {
	envVars = &EnvVars{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
	}
	if err := validateEnvs(); err != nil {
		return err
	}

	envVars.dsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		envVars.DBHost,
		envVars.DBPort,
		envVars.DBUser,
		envVars.DBPassword,
		envVars.DBName,
	)

	return nil
}

func validateEnvs() error {
	required := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	var missing []string
	for _, name := range required {
		value, exists := os.LookupEnv(name)
		if !exists || strings.TrimSpace(value) == "" {
			missing = append(missing, name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing or empty required environment variables: %s", strings.Join(missing, ", "))
	}

	return nil
}
