package main

import (
	"os"
	"strings"

	"github.com/m8b-dev/lukso-cli/config"
	"github.com/m8b-dev/lukso-cli/pid"
	"github.com/urfave/cli/v2"
)

const jwtSecretPath = configsRootDir + "/shared/secrets/jwt.hex"

// initializeDirectory initializes a working directory for lukso node, with all configurations for all networks
func initializeDirectory(ctx *cli.Context) error {
	if isAnyRunning() {
		return nil
	}

	if cfg.Exists() {
		message := "⚠️  This folder has already been initialized. Do you want to re-initialize it? Please note that configs in this folder will NOT be overwritten [y/N]:\n> "
		input := registerInputWithMessage(message)
		if !strings.EqualFold(input, "y") && input == "" {
			return nil
		}
	}

	for _, dependency := range clientDependencies {
		// this logic may fail when folder structure changes, but this shouldn't be the case
		if dependency.isBinary {
			continue
		}

		err := dependency.Download("", "", false, configPerms)
		if err != nil {
			log.Errorf("❌  There was error while downloading %s file: %v", dependency.name, err.Error())

			return nil
		}

	}

	err := createJwtSecret(jwtSecretPath)
	if err != nil {
		log.Errorf("❌  There was an error while creating JWT secret file: %v", err)

		return nil
	}

	err = os.MkdirAll(pid.FileDir, configPerms)
	if err != nil {
		log.Errorf("❌  There was an error while preparing PID directory: %v", err)

		return nil
	}

	log.Info("⚙️  Creating LUKSO configuration file...")

	err = cfg.Create("", "")
	if err != nil {
		log.Errorf("❌  There was an error while preparing LUKSO configuration: %v", err)

		return nil
	}

	log.Infof("✅  LUKSO configuration created under %s", config.Path)
	log.Info("✅  Folder initialized! \n1) ⚙️  Use 'lukso install' to install clients. \n2) ▶️  Use 'lukso start' to start your node.")

	return nil
}
