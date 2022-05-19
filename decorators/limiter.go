package decorators

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type Limiter struct {
	c.Decorator
	maxLoop int
}

func (l *Limiter) OnTick(t *c.Tick) util.Status {
	if l.GetChild() == nil {
		return util.ERROR
	}
	i := t.BlackBoard.GetInt("i", t.GetTree().GetID(), l.GetID())
	if i < l.maxLoop {
		status := l.GetChild().Execute(t)
		if status == util.SUCCESS || status == util.FAILURE {
			t.BlackBoard.Set("i", i+1, t.GetTree().GetID(), l.GetID())
		}
		return status
	}
	return util.FAILURE
}
