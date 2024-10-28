package common

import (
	"fmt"
	"github.com/magicsea/behavior3go/core"
	log "github.com/sirupsen/logrus"
)

func LogError(tick *core.Tick, err error) {
	if err == nil {
		return
	}

	log.Error(err)
	tick.Blackboard.SetMem(ErrorStr, err)
}

func LogErrorf(tick *core.Tick, format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)

	log.Error(err)
	tick.Blackboard.SetMem(ErrorStr, err)
}
