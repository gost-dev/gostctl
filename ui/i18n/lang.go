package i18n

import "sync"

type Lang struct {
	Name    Key
	Value   string
	content map[Key]string
}

var langs = []Lang{
	{
		Name:    English,
		Value:   "en_US",
		content: en_US,
	},
	{
		Name:    Chinese,
		Value:   "zh_CN",
		content: zh_CN,
	},
}

func Langs() []Lang {
	return langs
}

var (
	currentLang Lang = langs[0]
	mux         sync.RWMutex
)

func Current() Lang {
	mux.RLock()
	defer mux.RUnlock()

	return currentLang
}

func Set(lang string) {
	mux.Lock()
	defer mux.Unlock()

	for i := range langs {
		if langs[i].Value == lang {
			currentLang = langs[i]
			return
		}
	}
	currentLang = langs[0]
}

const (
	Server    Key = "server"
	Service   Key = "service"
	Chain     Key = "chain"
	Hop       Key = "hop"
	Auther    Key = "auther"
	Admission Key = "admission"
	Bypass    Key = "bypass"
	Resolver  Key = "resolver"
	Hosts     Key = "hosts"
	Limiter   Key = "limiter"
	Ingress   Key = "ingress"
	Observer  Key = "observer"
	Logger    Key = "logger"

	Type         Key = "type"
	Name         Key = "name"
	Address      Key = "address"
	Host         Key = "host"
	ServerName   Key = "serverName"
	Path         Key = "path"
	URL          Key = "url"
	URLHint      Key = "urlHint"
	Interval     Key = "interval"
	IntervalHint Key = "intervalHint"
	Timeout      Key = "timeout"
	TimeoutHint  Key = "timeoutHint"
	Seconds      Key = "seconds"
	Filter       Key = "filter"

	Auth      Key = "auth"
	BasicAuth Key = "basicAuth"
	Username  Key = "username"
	Password  Key = "password"

	Basic      Key = "basic"
	Advanced   Key = "advanced"
	AuthSimple Key = "authSimple"
	AuthAuther Key = "authAuther"

	Interface     Key = "interface"
	InterfaceHint Key = "interfaceHint"

	Handler           Key = "handler"
	Listener          Key = "listener"
	Forwarder         Key = "forwarder"
	Protocol          Key = "protocol"
	Nodes             Key = "nodes"
	Metadata          Key = "metadata"
	VerifyServerCert  Key = "verifyServerCert"
	RewriteHostHeader Key = "rewriteHostHeader"

	DirectoryPath  Key = "dirPath"
	CustomHostname Key = "customHostname"
	Hostname       Key = "hostname"
	EnableTLS      Key = "enableTLS"
	Keepalive      Key = "keepalive"
	DeleteServer   Key = "deleteServer"
	DeleteService  Key = "deleteService"

	ErrNameRequired Key = "errNameRequired"
	ErrNameExists   Key = "errNameExists"
	ErrURLRequired  Key = "errURLRequired"
	ErrInvalidAddr  Key = "errInvalidAddr"
	ErrDigitOnly    Key = "errDigitOnly"
	ErrDirectory    Key = "errDir"

	OK     Key = "ok"
	Cancel Key = "cancel"

	Running Key = "running"
	Ready   Key = "ready"
	Failed  Key = "failed"
	Closed  Key = "closed"
	Unknown Key = "unknown"

	Settings Key = "settings"
	Language Key = "language"
	English  Key = "english"
	Chinese  Key = "chinese"
	Theme    Key = "theme"
	Light    Key = "light"
	Dark     Key = "dark"
)

type Key string

func (k Key) Value() string {
	if k == "" {
		return ""
	}

	return get(k)
}

func get(key Key) string {
	mux.RLock()
	defer mux.RUnlock()

	if v := currentLang.content[key]; v != "" {
		return v
	}

	return langs[0].content[key]
}
