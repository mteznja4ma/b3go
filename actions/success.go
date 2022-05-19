package actions

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type Success struct {
	c.Action
}

func (s *Success) OnTick(t *c.Tick) util.Status {
	return util.SUCCESS
}
