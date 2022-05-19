package actions

import (
	c "github.com/mteznja4ma/Gai-b3Go/core"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

type Failure struct {
	c.Action
}

func (f *Failure) OnTick(t *c.Tick) util.Status {
	return util.FAILURE
}
