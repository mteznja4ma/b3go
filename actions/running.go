package actions

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type Running struct {
	c.Action
}

func (r *Running) OnTick(t *c.Tick) util.Status {
	return util.RUNNING
}
