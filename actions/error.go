package actions

import (
	c "github.com/mteznja4ma/Gai-b3Go/core"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

type Error struct {
	c.Action
}

func (e *Error) OnTick(tick *c.Tick) util.Status {
	return util.ERROR
}
