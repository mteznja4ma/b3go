package decorators

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type Inverter struct {
	c.Decorator
}

func (i *Inverter) OnTick(t *c.Tick) util.Status {
	if i.GetChild() == nil {
		return util.ERROR
	}
	status := i.GetChild().Execute(t)
	if status == util.SUCCESS {
		return util.FAILURE
	} else if status == util.FAILURE {
		return util.SUCCESS
	}
	return status
}
