package main

import (
	"fmt"
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

func readConfigDir(dir string) {
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
			readConfigDir(file)
		case mode.IsRegular(): //file
			ext := path.Ext(file) //file extension
			if ext != ".yml" && ext != ".yaml" {
				continue
			}
			ParseYML(file)
		}
	}
}

func getRandomHost() string {
	host := Group[0]
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
			if cacheHost(Group[0], v) {
				return v
			}
			if isHealth(myServ, v) == true {
				updateHostStatus(Group[0], v, fmt.Sprintf("%d", int32(time.Now().Unix())))
				return v
			}
			delete(myHost, hostIdentify)
			goto NEXT
		}
		cnt++
	}
	return host
}

func getWeightHost() string {
	host := Group[0]
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
				if cacheHost(Group[0], myHost[hostIdentify]) {
					return myHost[hostIdentify]
				}
				if isHealth(myServ, myHost[hostIdentify]) == true {
					updateHostStatus(Group[0], myHost[hostIdentify], fmt.Sprintf("%d", int32(time.Now().Unix())))
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

func getFoHost() string {
	host := Group[0]
	if s := len(myHost); s <= 0 {
		return host
	}
	for i := 0; i < len(myServ.order); i++ {
		hostIdentify = myServ.order[i]
		if cacheHost(Group[0], myHost[hostIdentify]) {
			return myHost[hostIdentify]
		}
		if isHealth(myServ, myHost[hostIdentify]) == true {
			updateHostStatus(Group[0], myHost[hostIdentify], fmt.Sprintf("%d", int32(time.Now().Unix())))
			return myHost[hostIdentify]
		}
	}
	return host
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
