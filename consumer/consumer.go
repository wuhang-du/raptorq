package consumer

import (
	"errors"
	"fmt"
	"wuhang-du/raptorq/common"
)

type Consumer struct {
	tracker      common.TrackerInterface
	Ck           *ChunkRecord
	pieceManager []*PieceManager
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

func (c *Consumer) Start(index int) error {
	//模拟单线程
	ck, err := NewChunkRecord(index)
	if err != nil {
		return err
	}
	//让我们荡起双桨
	c.Ck = ck

	for {
		if c.Ck.IsTimeout() {
			return errors.New("timeout")
		}

		if c.Ck.IsReady() {
			c.Ck.Close()
			ck, err := NewChunkRecord(c.Ck.Index + 1)
			if err != nil {
				return err
			}
			c.Ck = ck
		}

		// 看有没有节点过期，把arrange piece 给回来了；

		// 待分配piece
		for c.Ck.ArrangePiece > 0 {
			raq := c.tracker.GetRaqServer()
			raqChan := make(chan []common.PieceInfo, 1)
			err := raq.RegisterPiece("", 0, 4, raqChan)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			c.Ck.PieceCount -= 4
			pieceChan := NewPieceManager(c.Ck.Index, raqChan, "bkj")
			c.pieceManager = append(c.pieceManager, pieceChan)

			micro := c.tracker.GetMicroServer()
			microChan := make(chan []common.PieceInfo, 1)
			err = micro.MicroRegisterPiece("", 0, 4, microChan)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			pieceChan = NewPieceManager(c.Ck.Index, microChan, "skj")
			c.pieceManager = append(c.pieceManager, pieceChan)
		}

		for _, v := range c.pieceManager {
			//遍历每一个pieceChan
			data := v.GetData()
			if data != nil {
				for _, vv := range data {
					c.Ck.AddPiece(0, 1, vv.GetPiece())
				}
				continue
			}

			if v.IsTimeout() {
				// to do
				// delete piece count
			}
		}
	}
}
