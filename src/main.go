package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	gohost      string
	Group       []string
	DebugMode   bool
	noCache     bool
	healthCheck bool
	ConfigDir   string
	defDir      = "."
)

func init() {
	flag.BoolVar(&DebugMode, "d", false, "Debug mode")
	flag.BoolVar(&noCache, "nocache", false, "Cache mode [default Cached]")
	flag.BoolVar(&healthCheck, "check", false, "Health Check, will not get host [default false]")
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

func getGroupHost(group string, confdir ...string) string {
	var config string
	gohost = group

	if len(confdir) == 0 {
		config = ConfigDir
	} else {
		config = confdir[0]
	}
	if _, err := os.Stat(config); os.IsNotExist(err) {
		debug("Config not exist!(%s)\n", config)
		return gohost
	}

	readConfigDir(config, group)
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

func groupHealthCheck(group string, confdir ...string) {
	var config string
	if len(confdir) == 0 {
		config = ConfigDir
	} else {
		config = confdir[0]
	}
	if _, err := os.Stat(config); os.IsNotExist(err) {
		debug("Config not exist!(%s)\n", config)
	}
	readConfigDir(config, group)
	doHealthCheck(group)
}

func main() {
	flag.Parse()
	Group = flag.Args()
	if len(Group) != 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	debug("GROUP=>%v\n", Group)

	if healthCheck {
		groupHealthCheck(Group[0])
	} else {
		gohost = getGroupHost(Group[0])
		colorMsg(fmt.Sprintf("HOST(%s)=>%s\n", Group[0], gohost), color.FgHiGreen)
	}
}
