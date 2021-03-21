package microserver

import (
	"fmt"
	"wuhang-du/raptorq/common"
)

type MicroServer struct {
	tracker common.TrackerInterface
}

func NewMicroServer(tracker common.TrackerInterface) *MicroServer {
	m := &MicroServer{
		tracker: tracker,
	}

	tracker.RegisterMicroServer(m)
	return m
}

func (m *MicroServer) MicroRegisterPiece(uri string,
	startId int64, count int64, infoChan chan []common.PieceInfo) error {
	raqServer := m.tracker.GetRaqServer()
	raqChan := make(chan []common.PieceInfo, 1)
	err := raqServer.RegisterPiece(uri, startId, count, raqChan)
	if err != nil {
		return err
	}

	go func() {
		for {
			info := <-raqChan
			select {
			case infoChan <- info:
				fmt.Println("micro success proxy ", len(info))
			default:
				fmt.Println("micro infoChan error")
			}
		}
	}()

	return nil
}
