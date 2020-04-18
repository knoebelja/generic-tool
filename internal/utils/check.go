package utils

import (
	"fmt"
	"os"
)

// Catch checks for an error and exits
func Catch(err error) {
	Check(err, "")
}

// Check handles an error and suggests where to check to fix it before exiting
func Check(err error, msg string, values ...interface{}) {
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		if len(msg) > 0 {
			fmt.Println("CHECK:", fmt.Sprintf(msg, values...))
		}
		os.Exit(0)
	}
}
