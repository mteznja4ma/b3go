package core

import (
	"github.com/MTEzNjA4MA/Gai-b3Go/util"
)

type IBaseWrapper interface {
	execute(t *Tick) util.Status
	enter(t *Tick)
	open(t *Tick)
	tick(t *Tick) util.Status
	close(t *Tick)
	exit(t *Tick)
}

type IBaseNode interface {
	IBaseWrapper
	Ctor()
	Execute(tick *Tick) util.Status
	GetTitle() string
	GetCategory() string
	GetName() string
}

type BaseNode struct {
	IBaseWorker
	id          string
	name        string
	category    string
	title       string
	description string
}

func (b *BaseNode) SetName(name string) {
	b.name = name
}

func (b *BaseNode) SetTitle(title string) {
	b.title = title
}

func (b *BaseNode) SetDescription(description string) {
	b.description = description
}

func (b *BaseNode) Ctor() {}

func (b *BaseNode) GetID() string {
	return b.id
}

func (b *BaseNode) GetName() string {
	return b.name
}

func (b *BaseNode) GetTitle() string {
	return b.title
}

func (b *BaseNode) GetCategory() string {
	return b.category
}

func (b *BaseNode) Execute(t *Tick) util.Status {
	return b.execute(t)
}

func (b *BaseNode) execute(t *Tick) util.Status {
	// ENTER
	b.enter(t)

	// OPEN
	if t.BlackBoard.GetBool("isOpen", t.tree.id, b.id) {
		b.open(t)
	}

	// TICK
	status := b.tick(t)

	// CLOSE
	if status != util.RUNNING {
		b.close(t)
	}

	// EXIT
	b.exit(t)

	return status
}

func (b *BaseNode) enter(t *Tick) {
	t.enterNode(b)
	b.OnEnter(t)
}

func (b *BaseNode) open(t *Tick) {
	t.openNode(b)
	t.BlackBoard.Set("isOpen", true, t.tree.id, b.id)
	b.OnOpen(t)
}

func (b *BaseNode) tick(t *Tick) util.Status {
	t.tickNode(b)
	return b.OnTick(t)
}

func (b *BaseNode) close(t *Tick) {
	t.closeNode(b)
	t.BlackBoard.Set("isOpen", false, t.tree.id, b.id)
	b.OnClose(t)
}

func (b *BaseNode) exit(t *Tick) {
	t.exitNode(b)
	b.OnExit(t)
}
