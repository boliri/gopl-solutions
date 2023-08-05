package main

import (
	"fmt"
	"strconv"

	"eval"
)

func main() {
	var ex eval.Expr
	var vars = map[eval.Var]bool{}

	for {
		var e string

		fmt.Print("enter a mathematical expression (no spaces): ")
		_, err := fmt.Scanln(&e)
		if err != nil {
			fmt.Printf("\ninput error: %s\n", err)
			continue
		}

		ex, err = eval.Parse(e)
		if err != nil {
			fmt.Printf("parse error: %s\n", err)
			continue
		}

		fmt.Print("\nexpression parsed gracefully\n\n")

		err = ex.Check(vars)
		if err != nil {
			fmt.Printf("check error: %s\n\n", err)
			continue
		}

		fmt.Print("no check errors found in expression\n")
		break
	}

	env := eval.Env{}
	if len(vars) > 0 {
		fmt.Println("\nprovide values for the following variables:")
		for v := range vars {
			for {
				var val string

				fmt.Printf("\t-> %s = ", v)
				_, err := fmt.Scanln(&val)
				if err != nil {
					fmt.Printf("\ninput error: %s\n\n", err)
					continue
				}

				pval, err := strconv.ParseFloat(val, 64)
				if err != nil {
					fmt.Printf("value parse error: %s\n\n", err)
					continue
				}

				env[v] = pval
				break
			}
		}

		fmt.Print("\nall variable values provided gracefully\n")
	}

	fmt.Printf("\nexpression: %s\n", ex)
	fmt.Printf("environment: %v\n", env)
	fmt.Printf("evaluated expression: %.6g\n\n", ex.Eval(env))
}
