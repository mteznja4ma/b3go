package core

import "github.com/MTEzNjA4MA/Gai-b3Go/util"

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (c *Condition) Ctor() {

	c.category = util.CONDITION
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
/* func (this *Condition) Initialize(params *BTNodeCfg) {
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
} */
