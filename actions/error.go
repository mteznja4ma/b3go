package actions

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type Error struct {
	c.Action
}

func (e *Error) OnTick(tick *c.Tick) util.Status {
	return util.ERROR
}
