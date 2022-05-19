package actions

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type Failure struct {
	c.Action
}

func (f *Failure) OnTick(t *c.Tick) util.Status {
	return util.FAILURE
}
