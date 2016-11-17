package main

import "math/rand"
import "time"
import "github.com/kitech/php-go/phpgo"

func gfor_host(group string, conf string) string {
	return getGroupHost(group, conf)
}
func gfor_health(group string, conf string) {
	groupHealthCheck(group, conf)
}

func module_startup(ptype int, module_number int) int {
	//println("module_startup", ptype, module_number)
	return rand.Int()
}
func module_shutdown(ptype int, module_number int) int {
	//println("module_shutdown", ptype, module_number)
	return rand.Int()
}
func request_startup(ptype int, module_number int) int {
	//println("request_startup", ptype, module_number)
	return rand.Int()
}
func request_shutdown(ptype int, module_number int) int {
	//println("request_shutdown", ptype, module_number)
	return rand.Int()
}

func init() {
	//log.Println("run init...")
	rand.Seed(time.Now().UnixNano())

	phpgo.InitExtension("pg0", "")
	phpgo.RegisterInitFunctions(module_startup, module_shutdown, request_startup, request_shutdown)

	phpgo.AddFunc("gfor_host", gfor_host)
	phpgo.AddFunc("gfor_health", gfor_health)
}
