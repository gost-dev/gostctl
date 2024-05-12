package i18n

var zh_CN = map[Key]string{
	Server:    "服务器",
	Service:   "服务",
	Chain:     "转发链",
	Hop:       "跳跃点",
	Auther:    "认证器",
	Admission: "准入控制器",
	Bypass:    "分流器",
	Resolver:  "域名解析器",
	Hosts:     "域名映射器",
	Limiter:   "限速器",
	Ingress:   "Ingress",
	Observer:  "观测器",
	Logger:    "日志记录器",
	Recorder:  "数据记录器",
	Plugin:    "插件",
	Selector:  "选择器",
	Node:      "节点",

	Type:         "类型",
	Name:         "名称",
	Address:      "地址",
	Host:         "主机名",
	ServerName:   "服务名",
	Path:         "路径",
	URL:          "URL",
	URLHint:      "例如: http://localhost:8000",
	Interval:     "间隔时长",
	IntervalHint: "获取服务端配置的周期",
	Timeout:      "超时",
	TimeoutHint:  "获取配置时请求超时时间",
	Seconds:      "秒",
	Filter:       "过滤",
	TLS:          "TLS",
	HTTP:         "HTTP",
	CertFile:     "公钥文件",
	KeyFile:      "私钥文件",
	CAFile:       "CA证书",

	MetadataKey:   "键",
	MetadataValue: "值",

	Auth:          "认证",
	BasicAuth:     "认证",
	Username:      "用户名",
	Password:      "密码",
	Basic:         "基础",
	Advanced:      "高级",
	AuthSimple:    "单用户",
	AuthAuther:    "认证器",
	Interface:     "接口名(Interface)",
	InterfaceHint: "IP地址或接口名",

	FileServer:           "文件服务",
	SerialPortRedirector: "串口重定向",
	UnixDomainSocket:     "Unix Domain Socket",
	ReverseProxyTunnel:   "Tunnel",

	Handler:           "处理器(Handler)",
	Listener:          "监听器(Listener)",
	Forwarder:         "转发器(Forwarder)",
	Connector:         "连接器(Connector)",
	Dialer:            "拨号器(Dialer)",
	Protocol:          "协议",
	VerifyServerCert:  "验证服务端证书",
	Nodes:             "节点",
	Metadata:          "元数据",
	RewriteHostHeader: "重写Host头",

	DeleteServer:   "删除服务器？",
	DeleteService:  "删除服务？",
	DeleteChain:    "删除转发链？",
	DeleteHop:      "删除跳跃点？",
	DeleteNode:     "删除节点？",
	DeleteMetadata: "删除元数据？",

	SelectorStrategy: "策略",
	SelectorRound:    "轮询",
	SelectorRandom:   "随机",
	SelectorFIFO:     "主备",

	DataSource:       "数据源",
	FileDataSource:   "文件",
	FilePath:         "路径",
	RedisDataSource:  "Redis",
	RedisAddr:        "地址",
	RedisDB:          "数据库",
	RedisPassword:    "密码",
	RedisKey:         "Key",
	RedisType:        "类型",
	HTTPDataSource:   "HTTP",
	HTTPURL:          "URL",
	HTTPTimeout:      "请求超时(秒)",
	DataSourceReload: "重载周期",

	PluginGRPC: "gRPC",
	PluginHTTP: "HTTP",

	TimeSecond: "秒",

	DirectoryPath:  "文件目录路径",
	CustomHostname: "自定义主机名(重写HTTP Host头)",
	Hostname:       "主机名",
	EnableTLS:      "开启TLS",
	Keepalive:      "保持连接",

	ErrNameRequired: "名称必须填写",
	ErrNameExists:   "名称已存在",
	ErrURLRequired:  "URL必须填写",
	ErrInvalidAddr:  "无效的地址格式，仅支持[IP]:PORT或[HOST]:PORT",
	ErrDigitOnly:    "仅能输入数字",
	ErrDirectory:    "不是一个目录",

	OK:     "确认",
	Cancel: "取消",

	Running: "运行",
	Ready:   "就绪",
	Failed:  "失败",
	Closed:  "关闭",
	Unknown: "未知",

	Settings: "设置",
	Language: "语言",
	English:  "英语",
	Chinese:  "中文",
	Theme:    "主题",
	Light:    "浅色",
	Dark:     "深色",
}
