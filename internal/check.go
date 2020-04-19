package internal

import (
	"fmt"
	"os"
)

// catch checks for an error and exits
func catch(err error) {
	check(err, "")
}

// check handles an error and suggests where to check to fix it before exiting
func check(err error, msg string, values ...interface{}) {
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		if len(msg) > 0 {
			fmt.Println("CHECK:", fmt.Sprintf(msg, values...))
		}
		os.Exit(0)
	}
}

// out kills the app with a message why
func out(msg string, values ...interface{}) {
	fmt.Println("WATCH IT!", fmt.Sprintf(msg, values...))
	os.Exit(0)
}
