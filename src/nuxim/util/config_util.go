package util

import (
	"io/ioutil"
	"net/url"
	. "nuxim/defs/common"
	. "nuxim/defs/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func build_json_date_range(start, end string) string {
	msg := DateSelect{StartDate: start, EndDate: end}
	buf, err := msg.MarshalJSON()
	if err != nil {
		Log("Error in MarshalJSON date range!")
		return ""
	}
	return string(buf)
}

func build_date_range(data string) string {
	t := strings.Split(data, ",")
	if len(t) < 2 {
		k, err := strconv.Atoi(t[0])
		if err != nil {
			return build_json_date_range("-1", "0")
		}
		start := GetDate(k)
		return build_json_date_range(start, "0")
	}
	k, _ := strconv.Atoi(t[0])
	start := GetDate(k)
	k, _ = strconv.Atoi(t[1])
	end := GetDate(k)
	return build_json_date_range(start, end)
}

func build_id_list(data []string) string {
	msg := IDList{Data: data}
	buf, err := msg.MarshalJSON()
	if err != nil {
		Log("Error in MarshalJSON date range!")
		return ""
	}
	return string(buf)
}

func ParseParams(params string) url.Values {
	data := url.Values{}
	s := strings.Split(params, "&")
	for _, param := range s {
		h := strings.Split(param, "=")
		if len(h) < 2 {
			continue
		}
		k, v := h[0], h[1]
		t := strings.Split(v, ",")
		if k == "date_range" {
			data[k] = []string{build_date_range(v)}
			continue
		}
		switch {
		case len(t) > 1 && k != "cid":
			data[k] = []string{build_id_list(t)}
		default:
			data[k] = []string{v}
		}
	}
	return data
}

func ParseServiceParams(params string) (url.Values, url.Values) {
	control_params := url.Values{}
	s := strings.Split(params, "|")
	request_params := ParseParams(s[0])
	if len(s) > 1 {
		control_params = ParseParams(s[1])
	}
	return request_params, control_params
}

func FillDate(params url.Values, key, value string) {
	k, ok := params[key]
	if !ok {
		k = []string{value}
	}
	offset, _ := strconv.Atoi(k[0])
	params[key] = []string{GetDate(offset)}
}

func FillMissed(params url.Values, key, value string) {
	_, ok := params[key]
	if !ok {
		params[key] = []string{value}
	}
}

func FillMissedEmpty(params url.Values, key string) {
	FillMissed(params, key, "")
}

func FillRaw(params url.Values, key, value string, raw map[string]string) {
	k, ok := params[key]
	if ok {
		delete(params, key)
		raw[key] = k[0]
		return
	}
	raw[key] = value
}

func GetParam(params url.Values, key, value string) string {
	k, ok := params[key]
	if !ok {
		k = []string{value}
	}
	return k[0]
}

func GetParamInt(params url.Values, key, value string) int {
	k := GetParam(params, key, value)
	v, _ := strconv.Atoi(k)
	return v
}

func InitConfig(file string) Config {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		Log("Error in Reading Config File", err)
		os.Exit(1)
	}
	msg := Config{}
	err = msg.UnmarshalJSON(body)
	if err != nil {
		Log("Error in Parsing Config File", err)
		os.Exit(1)
	}
	return msg
}

func GetConfigPath(config Config, name string) string {
	v := config.Common
	p := v.Product
	s := v.Conf
	if name != COMPONENT_SDK && name != COMPONENT_SERVICES {
		switch name {
		case FOLDER_OUTPUT:
			s = v.Output
		case FOLDER_DATA:
			s = v.Data
		case FOLDER_LOGS:
			s = v.Logs
		case FOLDER_PIDS:
			s = v.Pids
		}
		return filepath.Join(s, p)
	}
	c := config.Services
	if name == COMPONENT_SDK {
		c = config.Sdk
	}
	return c.AutoLoading
}

func GetConfigs(path string) []string {
	s := []string{path, "/*.json"}
	h := strings.Join(s, "")
	list, err := filepath.Glob(h)
	if err != nil {
		Log("Error in Geting Config from", path, err)
		return []string{}
	}
	return list
}

func InitJobConfig(file string) CronJob {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		Log("Error in Reading Cron Job Config")
		os.Exit(1)
	}
	msg := CronJob{}
	err = msg.UnmarshalJSON(body)
	if err != nil {
		Log("Error in Parsing Cron Job Config")
		os.Exit(1)
	}
	return msg
}

func InitCronJobs(config Config, rootPath, compName string) []CronJob {
	s := GetConfigPath(config, compName)
	path := filepath.Join(rootPath, s)
	list := GetConfigs(path)
	jobs := []CronJob{}
	for _, file := range list {
		jobs = append(jobs, InitJobConfig(file))
	}
	return jobs
}
