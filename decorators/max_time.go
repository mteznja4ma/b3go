package decorators

import (
	"time"

	"github.com/mteznja4ma/Gai-b3Go/util"

	c "github.com/mteznja4ma/Gai-b3Go/core"
)

type MaxTime struct {
	c.Decorator
	maxTime int64
}

func (m *MaxTime) OnOpen(t *c.Tick) {
	startTime := time.Now().Unix()
	t.BlackBoard.Set("start_time", startTime, t.GetTree().GetID(), m.GetID())
}

func (m *MaxTime) OnTick(t *c.Tick) util.Status {
	if m.GetChild() == nil {
		return util.ERROR
	}
	currTime := time.Now().Unix()
	startTime := t.BlackBoard.GetInt64("startTime", t.GetTree().GetID(), m.GetID())
	if currTime-startTime > m.maxTime {
		return util.FAILURE
	}

	return m.GetChild().Execute(t)
}
