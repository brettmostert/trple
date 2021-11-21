package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("server is meant to do something")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
