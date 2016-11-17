package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	hostIdentify string
)

func readConfigDir(dir string, group string) {
	_, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	fio, _ := os.Stat(dir)
	if fio.Mode().IsDir() {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(0)
		}
		for _, f := range files {
			if fmt.Sprintf("%c", f.Name()[0]) == "." {
				continue
			}
			file := fmt.Sprintf("%s/%s", dir, f.Name())
			finfo, _ := os.Stat(file)
			switch mode := finfo.Mode(); {
			case mode.IsDir(): //directory
				defer readConfigDir(file, group)
			case mode.IsRegular(): //file
				ext := path.Ext(file) //file extension
				if ext != ".yml" && ext != ".yaml" {
					continue
				}
				ParseYML(file, group)
			}
		}
	} else if fio.Mode().IsRegular() {
		ext := path.Ext(dir) //file extension
		if ext != ".yml" && ext != ".yaml" {
			return
		}
		ParseYML(dir, group)
	}
}

func getRandomHost(group string) string {
	host := group
NEXT:
	s := len(myHost)
	if s <= 0 {
		return host
	}
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(s)
	cnt := 0
	for k, v := range myHost {
		if cnt == x {
			hostIdentify = k
			if cacheHost(group, v) {
				return v
			}
			if isHealth(myServ, v) == true {
				go updateHostStatus(group, v, fmt.Sprintf("%d", int32(time.Now().Unix())))
				return v
			}
			delete(myHost, hostIdentify)
			goto NEXT
		}
		cnt++
	}
	return host
}

func getWeightHost(group string) string {
	host := group
NEXT:
	s := len(myHost)
	if s <= 0 {
		return host
	}
	// get total weight
	totalweight := 0
	for _, v := range myWeight {
		totalweight += v
	}
	if totalweight > 0 {
		// get wight
		weight := 0
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(s)
		for k, v := range myWeight {
			weight += v
			if x < weight {
				hostIdentify = k
				if cacheHost(group, myHost[hostIdentify]) {
					return myHost[hostIdentify]
				}
				if isHealth(myServ, myHost[hostIdentify]) == true {
					go updateHostStatus(group, myHost[hostIdentify], fmt.Sprintf("%d", int32(time.Now().Unix())))
					return myHost[hostIdentify]
				}
				delete(myHost, hostIdentify)
				delete(myWeight, hostIdentify)
				goto NEXT
			}
		}
	}
	return host
}

func getFoHost(group string) string {
	host := group
	if s := len(myHost); s <= 0 {
		return host
	}
	for i := 0; i < len(myServ.order); i++ {
		hostIdentify = myServ.order[i]
		if cacheHost(group, myHost[hostIdentify]) {
			return myHost[hostIdentify]
		}
		if isHealth(myServ, myHost[hostIdentify]) == true {
			go updateHostStatus(group, myHost[hostIdentify], fmt.Sprintf("%d", int32(time.Now().Unix())))
			return myHost[hostIdentify]
		}
	}
	return host
}

func doHealthCheck(group string) {
	for node, host := range myHost {
		stime := time.Now().UnixNano() / int64(time.Millisecond)
		status := isHealth(myServ, host)
		etime := time.Now().UnixNano() / int64(time.Millisecond)
		if status {
			colorMsg(fmt.Sprintf("%dms\t%s\t%s\n", etime-stime, node, host), color.FgHiGreen)
			go updateHostStatus(group, host, fmt.Sprintf("%d", int32(time.Now().Unix())))
		} else {
			colorMsg(fmt.Sprintf("---\t%s\t%s\n", node, host), color.FgHiRed)
		}
	}
}

func cacheHost(group string, host string) bool {
	if noCache {
		return false
	}
	nowtime := int32(time.Now().Unix())
	cache, _ := strconv.ParseInt(getHostStatus(group, host), 10, 32)
	cachetime := int32(cache)
	if nowtime-cachetime > 600 { // host status cache time
		return false
	}
	debug("Cached!(%v)\n", group, host)
	return true
}
