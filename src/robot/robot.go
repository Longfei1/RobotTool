package robot

import "github.com/magicsea/behavior3go"

type IRobot interface {
	RegisterBevMap(mp *behavior3go.RegisterStructMaps)
}
