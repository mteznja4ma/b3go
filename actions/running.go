package actions

import (
	c "github.com/mteznja4ma/Gai-b3Go/core"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

type Running struct {
	c.Action
}

func (r *Running) OnTick(t *c.Tick) util.Status {
	return util.RUNNING
}
