// read input from the console:
package main

import (
	"bufio"
	"fmt"
	"os"
)

//`inputReader` 是一个指向 `bufio.Reader` 的指针。
var inputReader *bufio.Reader
var input string
var err error

func main() {
	//创建一个读取器，并将其与标准输入绑定。
	inputReader = bufio.NewReader(os.Stdin) // reader for input
	fmt.Println("Please enter some input: ")
	fmt.Println(inputReader.Size())
	input, err = inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
