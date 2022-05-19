package core

import (
	"github.com/mteznja4ma/b3go/config"
	"github.com/mteznja4ma/b3go/util"
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
	Category()
	Initialize(params *config.B3NodeConfig)
	Execute(tick *Tick) util.Status
	GetTitle() string
	GetCategory() string
	GetName() string
	SetBaseNodeWorker(w IBaseWorker)
	GetBaseNodeWorker() IBaseWorker
}

type BaseNode struct {
	IBaseWorker
	id          string
	name        string
	category    string
	title       string
	description string
	parameters  map[string]interface{}
	properties  map[string]interface{}
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

func (b *BaseNode) Category() {}

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

func (b *BaseNode) GetBaseNodeWorker() IBaseWorker {
	return b.IBaseWorker
}

func (b *BaseNode) SetBaseNodeWorker(w IBaseWorker) {
	b.IBaseWorker = w
}

func (b *BaseNode) Initialize(c *config.B3NodeConfig) {
	//this.id = b3.CreateUUID()
	//this.title       = this.title || this.name
	b.parameters = make(map[string]interface{})
	b.properties = make(map[string]interface{})

	b.id = c.Id //|| node.id;
	b.name = c.Name
	b.title = c.Title             //|| node.title;
	b.description = c.Description // || node.description;
	b.properties = c.Properties   //|| node.properties;

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
