package composites

import (
	c "github.com/mteznja4ma/Gai-b3Go/core"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

type Priority struct {
	c.Composite
}

func (p *Priority) OnTick(t *c.Tick) util.Status {
	for _, v := range p.GetAllChild() {
		status := v.Execute(t)
		if status != util.FAILURE {
			return status
		}
	}
	return util.FAILURE
}
