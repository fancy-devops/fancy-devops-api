package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize       int
	JwtSecret      string
	GitSecret      string
	GitAccessToken string
	FileBasePath   string
	RepositoryPath string
	BuildPath      string
	PublishPath    string
	ProductPath    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	GitSecret = sec.Key("GIT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	GitAccessToken = sec.Key("GIT_ACCESS_TOKEN").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	FileBasePath = sec.Key("FILE_BASEPATH").MustString("/home/fancy-devops")
	RepositoryPath = sec.Key("REPOSITORY_PATH").MustString("/repositories")
	BuildPath = sec.Key("BUILD_PATH").MustString("/build")
	PublishPath = sec.Key("PUBLISH_PATH").MustString("/publish")
	ProductPath = sec.Key("PRODUCT_PATH").MustString("/products")

}
