package main 


import (
	"github.com/golang-collections/collections/stack"
	valid "github.com/asaskevich/govalidator"
	"strconv"
	//"strings"
	)

func PolishNotation(arr []string) int {


	stack := stack.New()

	for _, value := range arr {

		if valid.IsInt(value) {
			number, _ := strconv.Atoi(value)
			stack.Push(number)
		} else {
			number_one := stack.Pop().(int)
			number_two := stack.Pop().(int)
		
			stack.Push(evaluate(number_two, number_one, value))
		}

	}

	return stack.Pop().(int)

}


func evaluate(number_one int, number_two int, oper string) int {

	switch oper {
	case "+":
		return number_one + number_two
	case "-":
		return number_one - number_two
	case "*":
		return number_one * number_two
		
	case "/":
		return number_one / number_two
		
	}
	return 0
}