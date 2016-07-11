package main

import "github.com/go-ini/ini"

var INI_LOCATION = `config.ini`
var API_PORT string


func init() {
	cfg, err := ini.Load(INI_LOCATION)
	if err != nil {
		logError(err, "init", "ini path: " + INI_LOCATION)
		return
	}

	API_PORT = cfg.Section("api").Key("port").String()
}