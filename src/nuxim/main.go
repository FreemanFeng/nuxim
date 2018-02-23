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
		"[-config config_file] [-cpuPort port] [-memPort port] [-cores cores_number]")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var configFile, memPort, cpuPort string
	var cores int
	var help bool

	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage

	flag.StringVar(&configFile, "config", DEFAULT_CONFIG_FILE, "config file")
	flag.StringVar(&memPort, "memPort", "", "port to profile memory")
	flag.StringVar(&cpuPort, "cpuPort", "", "port to profile cpu")
	flag.IntVar(&cores, "cores", 0, "cores to be used")

	flag.BoolVar(&help, "h", false, "Show Usage")
	flag.Parse()

	if help {
		Usage()
		return
	}

	util.ProfilingMemory(memPort)
	util.ProfilingCPU(cpuPort)
	util.UseCores(cores)

	config := util.InitConfig(configFile)
	rootPath := filepath.Dir(configFile)
	util.SavePID(config, rootPath, NUXIM_PID)

	quit := make(chan int)

	go sdk.Run(config, rootPath, quit)

	go services.Run(config, rootPath, quit)

	<-quit
}
