package main

import (
	"fmt"
	"github.com/fatih/color"
	"net"
	"net/http"
	"time"
)

func isHealth(myServ *Service, host string) bool {
	status := false
	debug(fmt.Sprintf("Check Server(%s) Health!\n", host), color.FgHiYellow)
	switch myServ.check {
	case "tcp":
		status = netCheck("tcp", host, myServ.port, myServ.timeout)
	case "udp":
		status = netCheck("udp", host, myServ.port, myServ.timeout)
	case "http":
		status = httpCheck("http", host, myServ.port, myServ.uri, myServ.timeout)
	case "https":
		status = httpCheck("https", host, myServ.port, myServ.uri, myServ.timeout)
	case "ping":
		status = icmpCheck(host)
	}
	return status
}

func netCheck(network string, host string, port int, timeout int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	_, err := net.DialTimeout(network, addr, time.Duration(timeout*1000000))
	if err != nil {
		colorMsg("FAIL: "+err.Error()+"\n", color.FgHiRed)
		return false
	}
	return true
}

func httpCheck(network string, host string, port int, uri string, timeout int) bool {
	url := fmt.Sprintf("%s://%s:%d%s", network, host, port, uri)
	if port == 0 {
		url = fmt.Sprintf("%s://%s%s", network, host, uri)
	}
	client := &http.Client{
		Timeout: time.Duration(timeout * 1000000),
	}
	_, err := client.Get(url)
	if err != nil {
		colorMsg("FAIL: "+err.Error()+"\n", color.FgHiRed)
		return false
	}
	return true
}

func icmpCheck(host string) bool {
	_, err := net.Dial("ip4:icmp", host)
	if err != nil {
		colorMsg("FAIL: "+err.Error()+"\n", color.FgHiRed)
		return false
	}
	return true
}
