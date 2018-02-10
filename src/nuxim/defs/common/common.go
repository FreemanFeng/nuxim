package common

// 全局配置
type Config struct {
	Sdk      ComponentConfig `json:"sdk"`      // 组件sdk的配置
	Services ComponentConfig `json:"services"` // 组件services的配置
	Common   CommonConfig    `json:"common"`   // 通用配置
}
type ComponentConfig struct {
	Port         int    `json:"port"`          // 监听端口
	ReadTimeout  int    `json:"read_timeout"`  // web服务读超时时间
	WriteTimeout int    `json:"write_timeout"` // web服务应答超时时间
	TaskTimeout  int    `json:"task_timeout"`  // 任务超时时间
	AutoLoading  string `json:"auto_loading"`  // 自动加载配置目录, i.e. services, 需拼接conf的路径
}
type CommonConfig struct {
	Product  string `json:"product"`   // 产品名称, nuxim
	Conf     string `json:"conf"`      // 配置目录相对于bin的路径, ../conf, 需拼接产品名称组成
	Output   string `json:"output"`    // 输出目录相对于bin的路径, ../output, 需拼接产品名称组成
	Data     string `json:"data"`      // 数据目录相对于bin的路径, ../data, 需拼接产品名称组成
	Pids     string `json:"pids"`      // 进程pid目录相对于bin的路径, ../pids, 需拼接产品名称组成
	Logs     string `json:"logs"`      // 日志目录相对于bin的路径, i.e. ../logs, 需拼接产品名称组成
	LogLevel string `json:"log_level"` // 日志level, 可选值: none/info/debug
	Online   int    `json:"online"`    // 是否是生产环境, 1是, 0否
}

// 调度任务配置
type CronJob struct {
	CronType int           `json:"cron_type"` // CRON类型
	Interval int           `json:"interval"`  // CRON JOB间隔
	Start    []int         `json:"start"`     // CRON JOB开始时间
	Platform string        `json:"platform"`  // 平台，SE
	Service  string        `json:"service"`   // 服务，STOCK
	Data     []ServiceData `json:"data"`      // 调用的服务列表
}
type ServiceData struct {
	ServiceID int    `json:"service_id"` // 服务ID，标识API接口
	Params    string `json:"params"`     // 服务参数, 其含义及具体字段由具体服务解析
}
