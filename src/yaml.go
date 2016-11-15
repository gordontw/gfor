package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Service struct {
	check   string // tcp|udp|http|https|ping
	method  string // random|weight|failover - default "random"
	uri     string
	port    int
	order   []string
	timeout int
}

var (
	myServ   = new(Service)
	myHost   map[string]string
	myWeight map[string]int
)

func ParseYML(yamlfile string) {
	yamlFile, _ := ioutil.ReadFile(yamlfile)
	any := map[string]interface{}{}
	err := yaml.Unmarshal(yamlFile, &any)
	if err != nil {
		log.Fatal(err)
	}

	myHost = make(map[string]string)
	myWeight = make(map[string]int)
	for k, v := range any {
		if k == Group[0] {
			flatten(k, v, myServ)
		}
	}
	debug("myServ=>%#v\n", myServ)
	debug("myHost=>%v\n", myHost)
	debug("myWeight=>%v\n", myWeight)
}

func flatten(prefix string, value interface{}, myServ *Service) {
	submap, ok := value.(map[interface{}]interface{})
	if ok {
		for k, v := range submap {
			flatten(prefix+"."+k.(string), v, myServ)
			switch k.(string) {
			case "check":
				myServ.check = v.(string)
			case "port":
				myServ.port = v.(int)
			case "method":
				myServ.method = v.(string)
			case "uri":
				myServ.method = v.(string)
			case "timeout":
				myServ.timeout = v.(int)
			case "order":
				orderAssign(v)
			case "host":
				nodeAssign(k.(string), v)
			case "weight":
				nodeAssign(k.(string), v)
			}
		}
	}
}

func orderAssign(value interface{}) {
	order, ok := value.([]interface{})
	if ok {
		for _, v := range order {
			myServ.order = append(myServ.order, v.(string))
		}
	}
}

func nodeAssign(t string, value interface{}) {
	//reflect.TypeOf(value)
	host, ok := value.(map[interface{}]interface{})
	if ok {
		for k, v := range host {
			if t == "host" {
				myHost[k.(string)] = v.(string)
			} else if t == "weight" {
				myWeight[k.(string)] = v.(int)
			}
		}
	}
}
