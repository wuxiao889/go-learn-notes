package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() {
	os.WriteFile(p.Title, p.Body, 0666)
}

func (p *Page) Load(title string) (err error) {
	p.Title = title
	p.Body, err = os.ReadFile(title)
	return err
}

func main() {
	p := &Page{
		"hello.md",
		[]byte("# hello\n## hello world"),
	}
	p.Save()
	var p1 Page
	p1.Load("hello.md")
	fmt.Print(p1)
}

func (p Page) String() string {
	return fmt.Sprintf("%s\n%s", p.Title, p.Body)
}
