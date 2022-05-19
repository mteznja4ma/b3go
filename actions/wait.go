package actions

import (
	"time"

	"github.com/mteznja4ma/b3go/config"
	"github.com/mteznja4ma/b3go/util"

	c "github.com/mteznja4ma/b3go/core"
)

type Wait struct {
	c.Action
	endTime int64
}

func (w *Wait) Initialize(setting *config.B3NodeConfig) {
	w.Action.Initialize(setting)
	w.endTime = setting.GetProperty2Int64("milliseconds")
}

func (w *Wait) OnOpen(t *c.Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	t.BlackBoard.Set("startTime", startTime, t.GetTree().GetID(), w.GetID())
}

func (w *Wait) OnTick(t *c.Tick) util.Status {
	var currTime int64 = time.Now().UnixNano() / 1000000
	var startTime = t.BlackBoard.GetInt64("startTime", t.GetTree().GetID(), w.GetID())
	if currTime-startTime > w.endTime {
		return util.SUCCESS
	}

	return util.RUNNING
}
