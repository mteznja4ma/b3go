package core

import "github.com/mteznja4ma/Gai-b3Go/util"

type SubTree struct {
	Action
}

/**
 *执行子树
 *使用sTree.Tick(tar, tick.Blackboard)的方法会导致每个树有自己的tick。
 *如果子树包含running状态，同时复用了子树会导致歧义。
 *改为只使用一个树，一个tick上下文。
**/
func (s *SubTree) OnTick(t *Tick) util.Status {

	//使用子树，必须先SetSubTreeLoadFunc
	//子树可能没有加载上来，所以要延迟加载执行
	sTree := subTreeLoadFunc(s.GetName())
	if nil == sTree {
		return util.ERROR
	}

	if t.GetTarget() == nil {
		panic("SubTree tick.GetTarget() nil !")
	}

	//tar := tick.GetTarget()
	//return sTree.Tick(tar, tick.Blackboard)

	t.pushSubtreeNode(s)
	ret := sTree.GetRoot().Execute(t)
	t.popSubtreeNode()
	return ret
}

func (this *SubTree) String() string {
	return "SBT_" + this.GetTitle()
}

var subTreeLoadFunc func(string) *BehaviorTree

//获取子树的方法
func SetSubTreeLoadFunc(f func(string) *BehaviorTree) {
	subTreeLoadFunc = f
}
