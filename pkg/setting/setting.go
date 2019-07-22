package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeOut time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fail to parse 'config/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("Run_Mode").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	RunMode = Cfg.Section("").Key("Run_Mode").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)

	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second

	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("fail to get section 'app':%v", err)
	}

	JwtSecret = sec.Key("JwtSecret").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PageSize").MustInt(10)
}
