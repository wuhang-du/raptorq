package consumer

import (
	"time"
	"wuhang-du/raptorq/common"
)

type PieceManager struct {
	nextIndex  int
	infoChan   chan []common.PieceInfo
	CreateTime int64
	raqtype    string
}

func NewPieceManager(index int, infoChan chan []common.PieceInfo, raqtype string) *PieceManager {
	return &PieceManager{
		nextIndex:  index,
		infoChan:   infoChan,
		CreateTime: time.Now().Unix(),
		raqtype:    raqtype,
	}
}

func (p *PieceManager) GetNextIndex() int {
	return p.nextIndex
}

func (p *PieceManager) GetData() []common.PieceInfo {
	select {
	case info := <-p.infoChan:
		return info
	default:
		return nil
	}
}

func (p *PieceManager) IsTimeout() bool {
	return false
}
