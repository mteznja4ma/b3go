package core

type Tick struct {
	tree       *BehaviorTree
	debug      interface{}
	target     interface{}
	BlackBoard *BlackBoard
	openNodes  []IBaseNode
	nodeCount  int
}

func NewTick() *Tick {
	t := &Tick{}
	t.Initialize()
	return t
}

func (t *Tick) Initialize() {
	t.tree = nil
	t.debug = nil
	t.target = nil
	t.BlackBoard = nil
	t.openNodes = make([]IBaseNode, 0)
	t.nodeCount = 0
}

func (t *Tick) GetTree() *BehaviorTree {
	return t.tree
}

func (t *Tick) GetTarget() interface{} {
	return t.target
}

func (t *Tick) enterNode(node IBaseNode) {
	t.nodeCount++
	t.openNodes = append(t.openNodes, node)
}

func (t *Tick) openNode(node *BaseNode) {

}

func (t *Tick) tickNode(node *BaseNode) {

}

func (t *Tick) closeNode(node *BaseNode) {
	nlen := len(t.openNodes)
	if nlen > 0 {
		t.openNodes = t.openNodes[:(nlen - 1)]
	}
}

func (t *Tick) exitNode(node *BaseNode) {

}
