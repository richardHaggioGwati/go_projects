package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	helloPrefix := prefix("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria

	//test for fileLen
	// if len(os.Args) < 2 {
	// 	log.Fatal("Can't find specified file || No file specified")
	// }
	// size, err := fileLen(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Fatal error: ", err)
	// 	return
	// }
	// fmt.Println("File size: ", size)
}

/*
Returning number of bytes in a file
*/

func fileLen(f string) (int, error) {
	// fileInfo, err := os.Stat(f)
	// if err != nil {
	// 	return 0, errors.New("Error opening file")
	// }
	// return int(fileInfo.Size()), nil

	// second approach using defer

	file, err := os.Open(f)
	if err != nil {
		return 0, errors.New("Error opening file")
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, errors.New("failed reading file metadata")
	}

	return int(fileInfo.Size()), nil
}

// exercise 1 simple calculator with err handling
func funcAsParametersMain() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
		}
		result, _ := opFunc(p1, p2)
		fmt.Println(result)
	}
}

func add(i, j int) (int, error) {
	return i + j, nil
}

func sub(i, j int) (int, error) {
	return i - j, nil
}

func mul(i, j int) (int, error) {
	return i * j, nil
}

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i + j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": sub,
}


//prefix function, exercise 
func prefix(s string) func(string) string {
	return func(t string) string {
		return s + " " + t
	}
}
