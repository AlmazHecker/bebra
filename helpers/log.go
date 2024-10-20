package helpers

import "fmt"

const (
    Reset  = "\033[0m"
    Yellow = "\033[33m"
	Red   = "\033[31m"  
)

func WarningLog(message string) {
	warningText := Yellow + "[WARNING]" + Reset + ": " + message
	fmt.Println(warningText)
}
func ErrorLog(message string) {
	warningText := Red + "[ERROR]" + Reset + ": " + message
	fmt.Println(warningText)
}