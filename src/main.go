package main

import (
	"flag"
	"fmt"
)

type Node struct {
	hostip string
	uri    string
	weight int
}

type Service struct {
	check   string
	method  string
	port    int
	timeout int
	nodes   []Node
}

var (
	group     []string
	debugMode bool
)

func init() {
	flag.BoolVar(&debugMode, "d", false, "Debug mode")
}

func debug(msg string, input ...interface{}) {
	if debugMode {
		fmt.Printf(msg, input)
	}
}

func main() {
	flag.Parse()
	group = flag.Args()

	debug("%v\n", group)
}
