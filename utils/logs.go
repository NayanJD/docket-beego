package utils

import (
	"github.com/beego/beego/v2/core/logs"
)

var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger()
	Log.SetLogger(logs.AdapterConsole)
	Log.Debug("logger set")
}
