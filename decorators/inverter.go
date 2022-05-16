package decorators

import (
	c "github.com/MTEzNjA4MA/Gai-b3Go/core"
	"github.com/MTEzNjA4MA/Gai-b3Go/util"
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
