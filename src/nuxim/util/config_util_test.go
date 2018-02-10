package util

import (
	"path/filepath"
	"testing"
)

func TestParseConfig(t *testing.T) {
	s := "a=b&c=d&e=f|h=i&j=k"
	p, c := ParseServiceParams(s)
	Log("req params:", p, "ctl params:", c)
}

func TestInitConfig(t *testing.T) {
	s := "/home/share/codes/go/conf/nuxim/nuxim.json"
	config := InitConfig(s)
	Log("req params:", config)
}

func TestGetConfigPath(t *testing.T) {
	s := "/home/share/codes/go/conf/nuxim/nuxim.json"
	config := InitConfig(s)
	path := filepath.Dir(s)
	h := GetConfigPath(config, "sdk")
	Log("sdk AutoLoading:", h, "root path", path)
	h = GetConfigPath(config, "services")
	Log("services AutoLoading:", h, "root path", path)
	h = GetConfigPath(config, "output")
	Log("output:", h, "root path", path)
	h = GetConfigPath(config, "data")
	Log("data:", h, "root path", path)
	h = GetConfigPath(config, "pids")
	Log("pids:", h, "root path", path)
	h = GetConfigPath(config, "logs")
	Log("logs:", h, "root path", path)
}
func TestInitCronJobs(t *testing.T) {
	s := "/home/share/codes/go/conf/nuxim/nuxim.json"
	config := InitConfig(s)
	rootPath := filepath.Dir(s)
	jobs := InitCronJobs(config, rootPath, "services")
	Log("Cron Jobs:", jobs)
}
