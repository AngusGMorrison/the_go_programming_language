// Write a program that reads a single expression from the standard input, prompts the user to
// provide values for any variables, then evaluates the expression in the resulting environment.
// Handle all errors gracefully.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"the_go_programming_language/ch7/cmd/ex15/eval"
)

func main() {
	fmt.Print("Enter an expression to evaluate: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	expr, err := eval.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	vars := make(map[eval.Var]bool)
	if err = expr.Check(vars); err != nil {
		log.Fatal(err)
	}
	env, err := getEnvVars(vars)
	if err != nil {
		log.Fatal(err)
	}
	result := expr.Eval(env)
	fmt.Printf("%s\t => %.6g\n", input, result)
}

func getEnvVars(vars map[eval.Var]bool) (eval.Env, error) {
	env := eval.Env{}
	s := bufio.NewScanner(os.Stdin)
	for k := range vars {
		for {
			fmt.Printf("Enter a value for %s: ", k)
			if !s.Scan() {
				return nil, fmt.Errorf("failed to scan input")
			} else if err := s.Err(); err != nil {
				return nil, err
			}

			val, err := strconv.ParseFloat(s.Text(), 64)
			if err != nil {
				fmt.Println("Invalid number. Please try again.")
			} else {
				env[k] = val
				break
			}

		}
	}
	return env, nil
}
