package raqserver

import (
	"wuhang-du/raptorq/common"

	"github.com/harmony-one/go-raptorq/pkg/defaults"
)

type RaqServer struct {
	commonOTI     uint64
	schemeSpecOTI uint32
}

func NewRaqServer(tracker common.TrackerInterface) *RaqServer {
	r := &RaqServer{}
	tracker.RegisterRaqServer(r)
	return r
}

func (r *RaqServer) GetRaqInfo() (uint64, uint32) {
	return r.commonOTI, r.schemeSpecOTI
}

func (r *RaqServer) RegisterPiece(uri string, id int64, count int64, infoChan chan []common.PieceInfo) error {
	//获取源文件
	info := `duwhabcd`
	//持续输出piece.
	_, err := defaults.NewEncoder([]byte(info), 1, 1, 8, 2)
	if err != nil {
		return err
	}

	return nil
}

func (r *RaqServer) MissData(uri string, id int64, count int64) ([]common.PieceInfo, error) {
	return nil, nil
}
