package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ask(question string) (answer string) {

	// 1. ask the question
	fmt.Println()
	fmt.Println(question)

	// 2. read the answer
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	catch(err)
	answer = strings.Replace(answer, "\n", "", 1)

	return
}

func askHint(question string, hints []string) (answer string) {

	switch len(hints) {
	case 0:
	case 1:
		question = question + "(hint: " + hints[0] + ")"
	default:
		question = question + " (hint: " + strings.Join(hints, ", ") + ", etc.)"
	}

	answer = ask(question)

	if len(answer) == 0 {
		fmt.Println("Please provide a response.")
		return askHint(question, []string{})
	}

	return

}

func askNum(question string) (answer float64) {

	answerString := ask(question)

	answer, err := strconv.ParseFloat(answerString, 64)

	if err != nil {
		if strings.Contains(err.Error(), "invalid syntax") {
			fmt.Println("Please provide a number.")
			return askNum(question)
		}
	}
	catch(err)

	return

}

func askYesNo(question string, defaultAnswer bool) (answer bool) {

	var questionWithDefault string
	switch defaultAnswer {
	case true:
		questionWithDefault = question + " (Y/n)"
	case false:
		questionWithDefault = question + " (y/N)"
	}

	answerString := ask(questionWithDefault)

	switch strings.ToUpper(answerString) {
	case "":
		return defaultAnswer
	case "YES", "Y":
		return true
	case "NO", "N":
		return false
	default:
		fmt.Println("Please provide a yes or no response.")
		return askYesNo(question, defaultAnswer)
	}
}
