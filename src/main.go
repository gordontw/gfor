package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	gohost    string
	Group     []string
	DebugMode bool
	ConfigDir string
	defDir    = "."
)

func init() {
	flag.BoolVar(&DebugMode, "d", false, "Debug mode")
	flag.StringVar(&ConfigDir, "c", defDir, "Debug mode")
}

func debug(msg string, input ...interface{}) {
	if DebugMode {
		fmt.Printf(msg, input)
	}
}

func colorMsg(msg string, c color.Attribute) {
	color.Set(c)
	fmt.Print(msg)
	color.Unset()
}

func main() {
	flag.Parse()
	Group = flag.Args()
	debug("GROUP=>%v\n", Group)
	if len(Group) > 1 {
		colorMsg("ERROR:too many args..\n", color.FgHiRed)
		os.Exit(1)
	}

	// work thru ConfigDir
	readConfigDir(ConfigDir)

	gohost = Group[0]
	switch myServ.method {
	case "random":
		gohost = getRandomHost()
	case "failover":
	case "weight":
		gohost = getWeightHost()
	default:
		gohost = getRandomHost()
	}

	fmt.Printf("HOST(%s)=>%s\n", Group[0], gohost)
}
