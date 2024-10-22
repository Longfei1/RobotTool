package api

import "RobotTool/src/jsdef/jsmsg"

type IAppApi interface {
	AddShowMsg(msg *jsmsg.ShowMessage)
	AddTipTag(tag *jsmsg.TipTag)
}
