package core

import (
	"fmt"

	"github.com/MTEzNjA4MA/Gai-b3Go/util"
)

type BehaviorTree struct {
	id          string
	title       string
	description string
	root        IBaseNode
	debug       interface{}
}

func NewBeTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (b *BehaviorTree) Initialize() {
	b.id = util.CreateUUID()
	b.title = "The behavior tree"
	b.description = "Default description"
	b.root = nil
}

func (b *BehaviorTree) GetID() string {
	return b.id
}

func (b *BehaviorTree) GetTitle() string {
	return b.title
}

func (b *BehaviorTree) GetDescription() string {
	return b.description
}

func (b *BehaviorTree) GetRoot() IBaseNode {
	return b.root
}

func (b *BehaviorTree) Tick(target interface{}, bb *BlackBoard) util.Status {
	if b == nil {
		panic("The blackboard parameter is obligatory and must be an instance of Blackboard")
	}

	/* CREATE A TICK OBJECT */
	var t = NewTick()
	t.debug = b.debug
	t.target = target
	t.BlackBoard = bb
	t.tree = b

	/* TICK NODE */
	var state = b.root.Execute(t)

	/* CLOSE NODES FROM LAST TICK, IF NEEDED */
	lastOpenNodes := bb.getTreeData(b.id).OpenNodes
	var curOpenNodes []IBaseNode
	curOpenNodes = append(curOpenNodes, t.openNodes...)

	// does not close if it is still open in this tick
	start := 0
	for i := 0; i < util.MinInt(len(lastOpenNodes), len(curOpenNodes)); i++ {
		start++
		if lastOpenNodes[i] != curOpenNodes[i] {
			break
		}
	}

	// close the nodes
	for i, v := range lastOpenNodes {
		if i >= start {
			v.close(t)
		}
	}

	/* POPULATE BLACKBOARD */
	bb.getTreeData(b.id).OpenNodes = curOpenNodes
	bb.SetTree("nodeCount", t.nodeCount, b.id)

	return state
}

func (b *BehaviorTree) Print() {
	printNode(b.root, 0)
}

func printNode(root IBaseNode, blk int) {
	for i := 0; i < blk; i++ {
		fmt.Print(" ")
	}

	fmt.Print("|-", root.GetTitle())

	if root.GetCategory() == util.DECORATOR {
		dec := root.(IDecorator)
		if dec.GetChild() != nil {
			printNode(dec.GetChild(), blk+3)
		}
	}

	fmt.Println("")

	if root.GetCategory() == util.COMPOSITE {
		comp := root.(IComposite)
		if comp.GetChildCount() > 0 {
			for i := 0; i < comp.GetChildCount(); i++ {
				printNode(comp.GetChild(i), blk+3)
			}
		}
	}
}
