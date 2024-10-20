package config

import (
	"bebra/helpers"
	"fmt"
	"os"
)

type Config struct {
	Apktool         string `json:"apktool"`
	Adb             string `json:"adb"`
	Signer          string `json:"signer"`
	DecompiledOutDir string `json:"decompiledOutDir"`
	CompiledOutDir  string `json:"compiledOutDir"`
}

func (c *Config) Validate() []string {
	var missing []string

	if c.Apktool == "" || !helpers.FileExists(c.Apktool){
		missing = append(missing, "apktool")
	}
	if c.Adb == "" || !helpers.FileExists(c.Adb){
		missing = append(missing, "adb")
	}
	if c.Signer == "" || !helpers.FileExists(c.Signer) {
		missing = append(missing, "apksigner")
	}

	return missing
}

func InitConfig(configPath string) Config {
	file, err := os.Open(configPath)
	if err != nil {
		helpers.ErrorLog(fmt.Sprintf("Error opening config file: %v\n", err))
		os.Exit(1)
	}
	defer file.Close()

	var config Config
	if err := helpers.JSONDecoder(file, &config); err != nil {
		helpers.ErrorLog("Error reading config file!")
		os.Exit(1)
	}

	missing := config.Validate()
	if len(missing) > 0 {
		helpers.WarningLog(fmt.Sprintf(
			"Some variables are missing or have incorrect locations in the config: %v\n"+
			"This is okay, but it may cause the program to work incorrectly.", missing))
	}

	if config.DecompiledOutDir == "" {
		config.DecompiledOutDir = "./decompiled"
	}
	if config.CompiledOutDir == "" {
		config.CompiledOutDir = "./build.apk"
	}

	return config
}

