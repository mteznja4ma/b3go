package composites

import (
	c "github.com/mteznja4ma/b3go/core"
	"github.com/mteznja4ma/b3go/util"
)

type MemSequence struct {
	c.Composite
}

func (m *MemSequence) OnOpen(t *c.Tick) {
	t.BlackBoard.Set("RunningChild", 0, t.GetTree().GetID(), m.GetID())
}

func (m *MemSequence) OnTick(t *c.Tick) util.Status {
	child := t.BlackBoard.GetInt("RunningChild", t.GetTree().GetID(), m.GetID())
	for i := child; i < m.GetChildCount(); i++ {
		status := m.GetChild(i).Execute(t)

		if status != util.SUCCESS {
			if status == util.RUNNING {
				t.BlackBoard.Set("RunningChild", i, t.GetTree().GetID(), m.GetID())
			}
			return status
		}
	}
	return util.SUCCESS
}
