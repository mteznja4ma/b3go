package core

import "github.com/mteznja4ma/b3go/util"

type IAction interface {
	IBaseNode
}

type Action struct {
	BaseNode
	BaseWorker
}

func (a *Action) Ctor() {
	a.category = util.ACTION
}
