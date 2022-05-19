package actions

import (
	c "github.com/mteznja4ma/Gai-b3Go/core"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

type Success struct {
	c.Action
}

func (s *Success) OnTick(t *c.Tick) util.Status {
	return util.SUCCESS
}
