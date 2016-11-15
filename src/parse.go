package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
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
	if s := len(myHost); s > 0 {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(s)
		cnt := 0
		for k, v := range myHost {
			if cnt == x {
				hostIdentify = k
				if isHealth(myServ, myHost[hostIdentify]) == true {
					host = v
					return host
				}
				delete(myHost, hostIdentify)
				goto NEXT
			}
			cnt++
		}
	}
	return host
}

func getWeightHost() string {
	host := Group[0]
NEXT:
	if s := len(myHost); s > 0 {
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
					if isHealth(myServ, myHost[hostIdentify]) == true {
						return myHost[hostIdentify]
					}
					delete(myHost, hostIdentify)
					delete(myWeight, hostIdentify)
					goto NEXT
				}
			}
		}
	}
	return host
}

func getFoHost() string {
	host := Group[0]
	if s := len(myHost); s > 0 {
		for i := 0; i < len(myServ.order); i++ {
			hostIdentify = myServ.order[i]
			if isHealth(myServ, myHost[hostIdentify]) == true {
				return myHost[myServ.order[i]]
			}
		}
	}
	return host
}
