package helpers

import (
	"fmt"
	"os"
	"strings"
)

func FileExists(filename string) bool {
    valid := checkPathPrefix(filename)
    if !valid {
        return false
    }

    info, err := isStatExist(filename)
    return err == nil && !info.IsDir() 
}

func DirExists(dirname string) bool {
    valid := checkPathPrefix(dirname)
    if !valid {
        return false
    }
    
    info, err := isStatExist(dirname)
    return err == nil && info.IsDir() 
}

func CreateFile(name string) *os.File {
    file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
    return file
}

func isStatExist(path string) (os.FileInfo, error) {
    info, err := os.Stat(path)
    if err != nil {
        return nil, err 
    }
    return info, nil 
}

func checkPathPrefix(path string) bool {
    // for linux
    if strings.HasPrefix(path, "~") {
        fmt.Printf("Replace '~' prefix with your home directory: %s\n", path)
        return false
    }
    return true
}
