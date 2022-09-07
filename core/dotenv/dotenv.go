package dotenv

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func SetVar(file string, name string, value string) error {
	content, loadErr := os.ReadFile(file)
	if loadErr != nil {
		log.WithField("category", "dotenv").Debugf("Error loading %s file: %s", file, loadErr.Error())
		content = []byte{}
	}

	env, readErr := godotenv.Unmarshal(string(content))
	if readErr != nil {
		log.WithField("category", "dotenv").Warnf("Error reading %s file: %s", file, loadErr.Error())
		return readErr
	}

	env[name] = value

	writeErr := godotenv.Write(env, file)
	if writeErr != nil {
		log.WithField("category", "dotenv").Errorf("Error reading %s file: %s", file, loadErr.Error())
		return writeErr
	}

	return nil
}

func DeleteVar(file string, name string) error {
	content, loadErr := os.ReadFile(file)
	if loadErr != nil {
		log.WithField("category", "dotenv").Debugf("Error loading %s file: %s", file, loadErr.Error())
		content = []byte{}
	}

	env, readErr := godotenv.Unmarshal(string(content))
	if readErr != nil {
		log.WithField("category", "dotenv").Warnf("Error reading %s file: %s", file, loadErr.Error())
		return readErr
	}

	delete(env, name)

	writeErr := godotenv.Write(env, file)
	if writeErr != nil {
		log.WithField("category", "dotenv").Errorf("Error reading %s file: %s", file, loadErr.Error())
		return writeErr
	}

	return nil
}
