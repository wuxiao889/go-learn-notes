package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('s')
	if err != nil {
		fmt.Println("error input")
		return
	}
	fmt.Printf("you input is %v\n", input)
}
