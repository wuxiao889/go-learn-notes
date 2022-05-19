package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var isNumbered = flag.Bool("n", false, "显示行号")

func cat(r *bufio.Reader) {
	for i := 0; ; i++ {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *isNumbered {
			fmt.Fprintf(os.Stdout, "%d:%s", i, buf)
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.PrintDefaults()
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
