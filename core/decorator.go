package core

import "github.com/MTEzNjA4MA/Gai-b3Go/util"

type IDecorator interface {
	IBaseNode
	SetChild(child IBaseNode)
	GetChild() IBaseNode
}

type Decorator struct {
	BaseNode
	BaseWorker
	child IBaseNode
}

func (d *Decorator) Ctor() {
	d.category = util.DECORATOR
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
/* func (b *Decorator) Initialize(params *BTNodeCfg) {
	b.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
} */

//GetChild
func (d *Decorator) GetChild() IBaseNode {
	return d.child
}

func (d *Decorator) SetChild(child IBaseNode) {
	d.child = child
}
