package consumer

import (
	"fmt"
	"wuhang-du/raptorq/common"
)

type Consumer struct {
	tracker common.TrackerInterface
}

/*
type Decoder struct {
	id     int64
	decode raptorq.Decoder
}
*/

func NewConsumer(tracker common.TrackerInterface) *Consumer {
	return &Consumer{tracker: tracker}
}

func (c *Consumer) Start() {
	raq := c.tracker.GetRaqServer()
	raqChan := make(chan []common.PieceInfo, 1)
	err := raq.RegisterPiece("", 0, 4, raqChan)
	if err != nil {
		fmt.Println(err.Error())
	}

	micro := c.tracker.GetMicroServer()
	microChan := make(chan []common.PieceInfo, 1)
	err = micro.MicroRegisterPiece("", 0, 4, microChan)
	if err != nil {
		fmt.Println(err.Error())
	}

}
