package q

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
)

func SampleQuestion() {
	menu := wmenu.NewMenu("What is your favorite food?")
	menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
	menu.Option("Pizza", nil, true, nil)
	menu.Option("Ice Cream", nil, false, nil)
	menu.Option("Tacos", nil, false, func(opt wmenu.Opt) error {
		fmt.Printf("Tacos are great")
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
