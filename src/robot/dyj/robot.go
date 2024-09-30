package dyj

import "github.com/magicsea/behavior3go"

type RobotDyj struct {
}

func (r *RobotDyj) RegisterBevMap(mp *behavior3go.RegisterStructMaps) {
	mp.Register("login", nil) //todo
}
