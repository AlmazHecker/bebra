package config

import (
	"bebra/helpers"
	"fmt"
	"os"
)


type Config struct {
	Apktool    string `json:"apktool,omitempty"`
	Adb        string `json:"adb,omitempty"`
	BuildTools string `json:"buildTools,omitempty"`
}

func GetConfig() Config {
	file, err := os.Open("bebra.config.json");
	if err != nil {
		panic(fmt.Errorf("Error opening file: %w", err))
	}
	defer file.Close()


	var config Config
	helpers.JSONDecoder(file, &config)

	var defaultConfig = getDefaultConfig()

	if !helpers.FileExists(config.Adb) {
		fmt.Printf("ADB not found, setting to default: %s\n", defaultConfig.Adb)
		config.Adb = defaultConfig.Adb
	}

	if !helpers.FileExists(config.Apktool) {
		fmt.Printf("APKTool not found, setting to default: %s\n", defaultConfig.Apktool)
		config.Apktool = defaultConfig.Apktool
	}

	if !helpers.DirExists(config.BuildTools) {
		fmt.Printf("Build Tools directory not found, setting to default: %s\n", defaultConfig.BuildTools)
		config.BuildTools = defaultConfig.BuildTools
	}

	return config
}

func getDefaultConfig() Config {
	homeDir,_ := os.UserHomeDir()
	Adb := "/usr/bin/adb";
	BuildTools := homeDir + "/Android/Sdk/build-tools/35.0.0";
	Apktool := "/usr/local/bin/apktool";

	return Config{ Adb: Adb, BuildTools: BuildTools, Apktool: Apktool }
}
// TODO no windows support for now
// if strings.Contains(runtime.GOOS, "windows") {
// 	adbDefaultPath = "C:\\Android\\platform-tools\\adb.exe"
// 	buildToolsPath = "C:\Android\platform-tools
// }