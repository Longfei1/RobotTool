package dyj

import "RobotTool/src/robot/dyj/pbgo"

func registerWsMsg(processor *MsgProcessor) {
	processor.Register(&pbgo.C2SLogin{}, uint16(pbgo.PCK_C2SLogin_ID))
	processor.Register(&pbgo.S2CLogin{}, uint16(pbgo.PCK_S2CLogin_ID))

	processor.Register(&pbgo.Ping{}, 10000)
	processor.Register(&pbgo.Pong{}, 10001)
}
