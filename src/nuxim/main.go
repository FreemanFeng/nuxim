package main

import (
	"flag"
	"fmt"
	"math"
	. "nuxim/defs/common"
	"nuxim/sdk"
	"nuxim/services"
	"nuxim/util"
	"os"
	"path/filepath"
	"strconv"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0],
		"[-cpuProfile filePath] [-memProfile port]")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var configFile, cpuProfile, memPort string
	var help bool

	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage

	flag.StringVar(&configFile, "config", DEFAULT_CONFIG_FILE, "config file")
	flag.StringVar(&cpuProfile, "cpuProfile", "", "write cpu profile to file")
	flag.StringVar(&memPort, "memProfile", "", "port to profile memory")

	flag.BoolVar(&help, "h", false, "Show Usage")
	flag.Parse()

	if help {
		Usage()
		return
	}

	util.ProfilingMemory(memPort)
	util.ProfilingCPU(cpuProfile)

	config := util.InitConfig(configFile)
	rootPath := filepath.Dir(configFile)
	util.SavePID(config, rootPath, NUXIM_PID)

	quit := make(chan int)

	go sdk.Run(config, rootPath, quit)

	go services.Run(config, rootPath, quit)

	<-quit
}
