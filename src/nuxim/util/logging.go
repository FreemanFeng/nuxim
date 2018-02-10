package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	. "nuxim/defs/common"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func SavePID(config Config, rootPath, pidName string) {
	pidPath := GetConfigPath(config, FOLDER_PIDS)
	path := filepath.Join(rootPath, pidPath)
	pid := os.Getpid()
	os.MkdirAll(path, os.ModeDir|os.ModePerm)
	s := []string{path, pidName}
	h := strings.Join(s, "/")
	s = []string{h, strconv.Itoa(pid)}
	h = strings.Join(s, ".")
	err := ioutil.WriteFile(h, []byte(strconv.Itoa(pid)), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func InitLogger(log_path, log_name string, ch chan *log.Logger) {
	os.MkdirAll(log_path, os.ModeDir|os.ModePerm)
	s := []string{log_path, log_name}
	h := strings.Join(s, "/")
	logfile, err := os.Create(h)
	defer logfile.Close()
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, logfile)

	logger := log.New(logfile, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.SetOutput(mw)
	ch <- logger

	quit := make(chan int)
	// 阻塞直到程序退出
	<-quit
}

func Log(params ...interface{}) {
	fmt.Print(Now(), " ")
	fmt.Println(params...)
}
