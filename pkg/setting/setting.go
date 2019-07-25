package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)

	if err != nil {
		log.Fatalf("Cfg.MapTo Appsetting err:%v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}

// func init() {
// 	var err error
// 	Cfg, err = ini.Load("conf/app.ini")
// 	if err != nil {
// 		log.Fatalf("fail to parse 'config/app.ini': %v", err)
// 	}

// 	LoadBase()
// 	LoadServer()
// 	LoadApp()
// }

// func LoadBase() {
// 	RunMode = Cfg.Section("").Key("Run_Mode").MustString("debug")
// }

// func LoadServer() {
// 	sec, err := Cfg.GetSection("server")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'server':%v", err)
// 	}
// 	RunMode = Cfg.Section("").Key("Run_Mode").MustString("debug")

// 	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)

// 	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second

// 	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
// }

// func LoadApp() {
// 	sec, err := Cfg.GetSection("app")
// 	if err != nil {
// 		log.Fatalf("fail to get section 'app':%v", err)
// 	}

// 	JwtSecret = sec.Key("JwtSecret").MustString("!@)*#)!@U#@*!@!)")
// 	PageSize = sec.Key("PageSize").MustInt(10)
// 	ExpireTime = sec.Key("ExpireTime").MustInt(3)
// }
