package decorators

import (
	c "github.com/MTEzNjA4MA/Gai-b3Go/core"
	"github.com/MTEzNjA4MA/Gai-b3Go/util"
)

type Repeater struct {
	c.Decorator
	maxLoop int
}

func (r *Repeater) OnOpen(t *c.Tick) {
	t.BlackBoard.Set("i", 0, t.GetTree().GetID(), r.GetID())
}

func (r *Repeater) OnTick(t *c.Tick) util.Status {
	if r.GetChild() == nil {
		return util.ERROR
	}
	i := t.BlackBoard.GetInt("i", t.GetTree().GetID(), r.GetID())
	status := util.ERROR
	for r.maxLoop < 0 || i < r.maxLoop {
		status = r.GetChild().Execute(t)
		if status == util.FAILURE || status == util.SUCCESS {
			i++
		} else {
			break
		}
	}

	t.BlackBoard.Set("i", i, t.GetTree().GetID(), r.GetID())
	return status
}
