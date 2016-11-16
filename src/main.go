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
	noCache   bool
	ConfigDir string
	defDir    = "."
)

func init() {
	flag.BoolVar(&DebugMode, "d", false, "Debug mode")
	flag.BoolVar(&noCache, "nocache", false, "Cache mode [default Cached]")
	flag.StringVar(&ConfigDir, "c", defDir, "YAML directory")
}

func debug(msg string, input ...interface{}) {
	if DebugMode {
		fmt.Printf("DEBUG: "+msg, input)
	}
}

func colorMsg(msg string, c color.Attribute) {
	color.Set(c)
	fmt.Print(msg)
	color.Unset()
}

func getGroupHost(group string) string {
	readConfigDir(ConfigDir, group)

	gohost = group
	switch myServ.method {
	case "random":
		gohost = getRandomHost(group)
	case "failover":
		gohost = getFoHost(group)
	case "weight":
		gohost = getWeightHost(group)
	default:
		gohost = getRandomHost(group)
	}
	return gohost
}

func main() {
	flag.Parse()
	Group = flag.Args()
	if len(Group) == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if len(Group) > 1 {
		colorMsg("ERROR:too many args..\n", color.FgHiRed)
		os.Exit(0)
	}
	debug("GROUP=>%v\n", Group)

	gohost = getGroupHost(Group[0])
	colorMsg(fmt.Sprintf("HOST(%s)=>%s\n", Group[0], gohost), color.FgHiGreen)
}
