package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

type registration struct {
	Name      string  `json:"name" yaml:"name"`
	Age       float64 `json:"age" yaml:"age"`
	Boy       bool    `json:"boy" yaml:"boy"`
	Sport     string  `json:"sport" yaml:"sport"`
	GPA       float64 `json:"gpa" yaml:"gpa"`
	Volunteer string  `json:"volunteer" yaml:"volunteer"`
}

// New asks a question
func New(cmd *cobra.Command, opts []string) {

	fmt.Println("-- REGISTER -- ALL BOYS SCHOOL OF AMERICA --")
	fmt.Println("--------------------------------------------")

	name := askHint("What's your name?", []string{"it's on your birth certificate"})

	br()

	age := askNum("How old are you?")
	br()

	gpa := askNum("What's your GPA?")
	br()

	var form *registration
	readYAML("form.yaml", &form)

	form.Name = name
	form.Age = age
	form.GPA = gpa

	writeYAML(name+".yaml", &form)

	fmt.Println("Thank you for submission")

}

func br() {
	fmt.Println()
	fmt.Println("--------------------------------------------")
}
