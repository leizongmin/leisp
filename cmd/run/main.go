package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leizongmin/leisp/interpreter"
)

func main() {
	flag.Parse()

	files := flag.Args()
	for _, f := range files {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		r := interpreter.Eval(nil, string(content))
		if r.IsError() {
			fmt.Println(r.ToString())
			os.Exit(2)
		}
	}
}
