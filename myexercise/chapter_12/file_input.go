package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./source.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	inputReader := bufio.NewReader(file)
	for {
		s, err := inputReader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("complete")
			break
		}
		sa := strings.Split(s[:len(s)-2], "|")
		price, err := strconv.Atoi(sa[2])
		if err != nil {
			fmt.Println(err)
		}
		p := product{sa[0], sa[1], price}
		fmt.Println(p)
	}

}

type product struct {
	title    string
	fullname string
	price    int
}

func (p product) String() string {
	return fmt.Sprintf("title:%s fullname:%s price:%d\n", p.title, p.fullname, p.price)
}
