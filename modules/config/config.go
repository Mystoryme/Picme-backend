package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	mod "picme-backend/modules"
	"picme-backend/types/module"
	"picme-backend/utils/text"
)

func Init() {
	// Initialize blank configuration struct
	conf := new(module.Config)

	// Load configurations to struct
	yml, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Fatal("UNABLE TO READ YAML CONFIGURATION FILE")
	}
	if err := yaml.Unmarshal(yml, conf); err != nil {
		logrus.Fatal("UNABLE TO PARSE YAML CONFIGURATION FILE")
	}

	// Validate configurations
	if err := text.Validator.Struct(conf); err != nil {
		logrus.Fatal("INVALID CONFIGURATION: " + err.Error())
	}

	// Apply log level configuration
	spew.Config = spew.ConfigState{Indent: "  "}

	// * Assign to module
	mod.Conf = conf
}
