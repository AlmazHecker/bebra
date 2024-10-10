package helpers

import (
	"encoding/json"
	"fmt"
	"io"
)

func JSONDecoder(reader io.Reader, v interface{}) error {
    decoder := json.NewDecoder(reader)
    if err := decoder.Decode(v); err != nil {
        return fmt.Errorf("error decoding JSON: %v", err)
    }
    return nil
}