package decorators

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type RepeatUntilFailure struct {
	c.Decorator
	maxLoop int
}

func (r *RepeatUntilFailure) OnOpen(t *c.Tick) {
	t.BlackBoard.Set("i", 0, t.GetTree().GetID(), r.GetID())
}

func (r *RepeatUntilFailure) OnTick(t *c.Tick) util.Status {
	if r.GetChild() == nil {
		return util.ERROR
	}
	i := t.BlackBoard.GetInt("i", t.GetTree().GetID(), r.GetID())
	status := util.ERROR
	for r.maxLoop < 0 || i < r.maxLoop {
		status = r.GetChild().Execute(t)
		if status == util.SUCCESS {
			i++
		} else {
			break
		}
	}

	t.BlackBoard.Set("i", i, t.GetTree().GetID(), r.GetID())
	return status
}
