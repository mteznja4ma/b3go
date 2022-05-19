package composites

import (
	c "github.com/mteznja4ma/Gai-b3Go/core"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

type Sequence struct {
	c.Composite
}

func (s *Sequence) Tick(t *c.Tick) util.Status {
	for _, v := range s.GetAllChild() {
		status := v.Execute(t)
		if status != util.SUCCESS {
			return status
		}
	}
	return util.SUCCESS
}
