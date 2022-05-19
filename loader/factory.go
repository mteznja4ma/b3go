package loader

import (
	a "github.com/mteznja4ma/Gai-b3Go/actions"
	c "github.com/mteznja4ma/Gai-b3Go/composites"
	"github.com/mteznja4ma/Gai-b3Go/config"
	"github.com/mteznja4ma/Gai-b3Go/core"
	d "github.com/mteznja4ma/Gai-b3Go/decorators"
	"github.com/mteznja4ma/Gai-b3Go/util"
)

func createBaseCategory() *util.RegisterStructMaps {

	s := util.NewRegisterStructMaps()
	//Actions
	s.Register("Error", &a.Error{})
	s.Register("Failure", &a.Failure{})
	s.Register("Running", &a.Running{})
	s.Register("Success", &a.Success{})
	s.Register("Wait", &a.Wait{})

	//Composites
	s.Register("MemPriority", &c.MemPriority{})
	s.Register("MemSequence", &c.MemSequence{})
	s.Register("Priority", &c.Priority{})
	s.Register("Sequence", &c.Sequence{})

	//Decorators
	s.Register("Inverter", &d.Inverter{})
	s.Register("Limiter", &d.Limiter{})
	s.Register("MaxTime", &d.MaxTime{})
	s.Register("Repeater", &d.Repeater{})
	s.Register("RepeatUntilFailure", &d.RepeatUntilFailure{})
	s.Register("RepeatUntilSuccess", &d.RepeatUntilSuccess{})

	return s
}

func CreateB3Tree(c *config.B3TreeConfig, extra *util.RegisterStructMaps) *core.BehaviorTree {
	b := createBaseCategory()
	t := core.NewBeTree()
	t.Load(c, b, extra)
	return t
}
