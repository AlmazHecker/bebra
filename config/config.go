package config

import (
	"bebra/helpers"
	"fmt"
	"os"
)


type Config struct {
	Apktool    string `json:"apktool,omitempty"`
	Adb        string `json:"adb,omitempty"`
	BuildTools string `json:"buildTools"`
}

func GetConfig() {
	file, err := os.Open("bebra.config.json");
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var config Config
	if err := helpers.JSONDecoder(file, &config); err != nil {
		fmt.Println(err)
	}

	
	fmt.Println("DECODED JSON", config.Adb)

    // // Decode the JSON data into the struct
	// decoder := json.NewDecoder(file)
	// if err := decoder.Decode(&config); err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }



	fmt.Printf("apktool: %s\nAdb: %s\nbuildtools: %s\n", config.Apktool, config.Adb, config.BuildTools)
}