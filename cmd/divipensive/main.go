// Package main provides a cli for divipensive
package main

import (
	"fmt"

	"github.com/kitschysynq/divipensive"
)

func main() {
	for _, t := range divipensive.Tolerances() {
		fmt.Println("Tolerance family: ", divipensive.ToleranceString(t))
		for _, d := range divipensive.Decades() {
			for i := 0; i < int(t); i++ {
				fmt.Println(divipensive.StringValue(i, t, d))
			}
		}
	}
}
