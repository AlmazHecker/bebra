package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func JSONDecoder(reader io.Reader, v interface{}) error {
    decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
    if err := decoder.Decode(v); err != nil {
        fmt.Printf("error decoding JSON: %v\n", err)
		os.Exit(1)
    }
    return nil
}

func JSONEncoder(file *os.File, v interface{}) {
    encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Indent the JSON for readability
	err := encoder.Encode(v)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}