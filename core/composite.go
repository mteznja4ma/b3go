package core

import (
	"fmt"

	"github.com/mteznja4ma/Gai-b3Go/util"
)

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	BaseWorker

	children []IBaseNode
}

func (c *Composite) Ctor() {

	c.category = util.COMPOSITE
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
/* func (this *Composite) Initialize(params *BTNodeCfg) {
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.children = make([]IBaseNode, 0)
	//fmt.Println("Composite Initialize")
} */

/**
 *
 * @method GetChildCount
 * @getChildCount
**/
func (c *Composite) GetChildCount() int {
	return len(c.children)
}

//GetChild
func (c *Composite) GetChild(index int) IBaseNode {
	return c.children[index]
}

func (c *Composite) GetAllChild() []IBaseNode {
	return c.children
}

//AddChild
func (c *Composite) AddChild(child IBaseNode) {
	c.children = append(c.children, child)
}
func (c *Composite) tick(tick *Tick) util.Status {
	fmt.Println("tick Composite1")
	return util.ERROR
}
