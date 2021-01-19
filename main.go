package main

import (
	"fmt"
	"os"
	"strconv"
)

func checknum() {
	var num int64
	for {
		fmt.Print("Enter Decimal Number:")
		fmt.Scanln(&num)
		if num >= 0 {
			output := strconv.FormatInt(num, 2)
			fmt.Print("Output: ", output)
			os.Exit(0)
		} else {
			fmt.Println("invalid input")
		}

	}
}

func main() {
	for {
		fmt.Println("Press 1 to run")
		fmt.Println("Press 2 to exit")
		var input int
		fmt.Scanln(&input)
		if input < 1 {
			fmt.Println("invalid input")
		} else if input > 2 {
			fmt.Println("invalid input")
		} else if input == 1 {
			checknum()
		} else {
			break
		}
	}

}
