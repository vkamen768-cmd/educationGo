package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a float64, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("undefined operation")
	}
}

func getInput() (string, error) {
	//var a float64
	//var b float64
	//var op string
	fmt.Print("Enter: a op b ? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("invalid input")
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", errors.New("invalid input count zero")
	}
	return input, nil

}

func printResult(res float64) {
	fmt.Printf("Result: %.2f\n", res)

}

func askContinue() bool {
	fmt.Print("Continue? (y/n): ")
	var input string
	fmt.Scan(&input)
	return strings.ToLower(input) == "y"
}

func main() {
	var history []float64
	for {

		input, err := getInput()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		if input == "history" {
			for k, v := range history {
				fmt.Printf("%d: %.2f\n", k+1, v)
			}
			continue
		}
		parts := strings.Split(input, " ")

		if len(parts) != 3 {
			fmt.Println("invalid input")
			continue
		}

		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("invalid input a")
			continue
		}
		b, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("invalid input b")
			continue
		}
		op := parts[1]
		res, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		printResult(res)
		history = append(history, res)
		if !askContinue() {
			break
		}
	}

}
