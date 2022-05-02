package core

import (
	"github.com/MTEzNjA4MA/Gai-b3Go/util"
)

type IBaseWorker interface {
	OnEnter(t *Tick)

	OnOpen(t *Tick)

	OnTick(t *Tick) util.Status

	OnClose(t *Tick)

	OnExit(t *Tick)
}

type BaseWorker struct{}

func (b *BaseWorker) OnEnter(t *Tick) {

}

func (b *BaseWorker) OnOpen(t *Tick) {

}

func (b *BaseWorker) OnTick(t *Tick) util.Status {
	return util.ERROR
}

func (b *BaseWorker) OnClose(t *Tick) {

}

func (b *BaseWorker) OnExit(t *Tick) {

}
