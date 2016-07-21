package main

import (
	"net/http"
	_ "1million/tools"
	_ "1million/routers"
	"1million/tools"
	_ "net/http/pprof"
	"runtime"
	
//	"log"
)

func main() {
	var MULTICORE int = runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(MULTICORE)
	post := tools.Conf.GetValue("default", "httpport")
	http.ListenAndServe("10.157.193.53:" + post, nil)
}


