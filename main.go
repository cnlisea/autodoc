package main

import (
	"github.com/cnlisea/automation"
	"github.com/cnlisea/automation/config"
	"flag"
	"fmt"
	"strings"
	"github.com/astaxie/beego/logs"
)

var configFile *string = flag.String("config", "/etc/am.json", "automation config file")
var logLevel *string = flag.String("log-level", "debug", "log level [debug|info|warn|error], default debug")
var version *bool = flag.Bool("v", false, "the version of automation")
var baseUrl *string = flag.String("u", "http://test.jubao56.com", "automation base url")

//var authType *string = flag.String("a", "unauth", "auth type [unauth|bearer], default unauth")



const banner string = `
automation v1.0.1
`
// Automation
var _ = ``

func main() {
	fmt.Print(banner)
	flag.Parse()
	if *version {
		return
	}

	if len(*configFile) == 0 {
		logs.Error("must use a config file")
		return
	}

	if len(*baseUrl) != 0 {
		config.BaseUrl = *baseUrl
	}

	// logs settings
	logs.EnableFuncCallDepth(true)
	// logs output filename and logs level default
	if "" != *logLevel {
		setLogLevel(*logLevel)
	}

	cfg, err := automation.ParseConfigFile(*configFile)
	if nil != err {
		logs.Error("parse config fail, err: ", err)
		return
	}

	instance := automation.New(cfg)
	if err := instance.Parse(); nil != err {
		logs.Error("automation parse err:", err)
		return
	}

	if err := instance.Run(); nil != err {
		logs.Error("automation run err:", err)
		return
	}

	logs.Info("instance successfully!!!")
}


func setLogLevel(level string) {
	// set default level info
	var l int = logs.LevelDebug
	switch strings.ToLower(level) {
	case "debug":
		l = logs.LevelDebug
	case "info":
		l = logs.LevelInformational
	case "warn":
		l = logs.LevelWarning
	case "error":
		l = logs.LevelError
	}
	logs.SetLevel(l)
}