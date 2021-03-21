package tracker

import "wuhang-du/raptorq/common"

type Tracker struct {
	raq   common.RaqServerInterface
	micro common.MicroInterface
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) RegisterRaqServer(r common.RaqServerInterface) {
	t.raq = r
}

func (t *Tracker) RegisterMicroServer(m common.MicroInterface) {
	t.micro = m
}

func (t *Tracker) GetRaqServer() common.RaqServerInterface {
	return t.raq
}

func (t *Tracker) GetMicroServer() common.MicroInterface {
	return t.micro
}
