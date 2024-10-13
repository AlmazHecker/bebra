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
	DecompileOutDir string `json:"decompileOutDir,omitempty"`
}

func GetConfig(configPath string) Config {
	file, err := os.Open(configPath);
	if err != nil {
		fmt.Println("Error opening file: %w", err)
		os.Exit(1)
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

	if config.DecompileOutDir == "" {
		config.DecompileOutDir = defaultConfig.DecompileOutDir
	}

	return config
}

func getDefaultConfig() Config {
	homeDir,_ := os.UserHomeDir()
	Adb := "/usr/bin/adb";
	BuildTools := homeDir + "/Android/Sdk/build-tools/35.0.0";
	Apktool := "/usr/local/bin/apktool";
	DecompileOutDir := "./decompiled"

	return Config{ Adb: Adb, BuildTools: BuildTools, Apktool: Apktool,DecompileOutDir: DecompileOutDir }
}

// TODO no windows support for now
// if strings.Contains(runtime.GOOS, "windows") {
// 	adbDefaultPath = "C:\\Android\\platform-tools\\adb.exe"
// 	buildToolsPath = "C:\Android\platform-tools
// }