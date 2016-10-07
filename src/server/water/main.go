package main

import (
	"flag"
	"fmt"
	"runtime"

	log "github.com/dongzerun/logger"

	"conf"
	"http"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	config_file = flag.String("config_file", "", "config file abs path")
)

func main() {
	flag.Parse()
	fmt.Println("config_file: ", *config_file)
	if *config_file == "" {
		fmt.Println("config_file must not empty")
		return
	}

	conf.InitConfig(*config_file)

	init_logger()

	http.Start()
}

func init_logger() {
	log.InitLooger(
		conf.GlobalConfig.LogConfig.Level,
		conf.GlobalConfig.LogConfig.Dir,
		conf.GlobalConfig.LogConfig.Filename,
		conf.GlobalConfig.LogConfig.ReserveNum,
		conf.GlobalConfig.LogConfig.Suffix,
		conf.GlobalConfig.LogConfig.Console,
		conf.GlobalConfig.LogConfig.Colorfull)
}
