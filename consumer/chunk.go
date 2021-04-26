package consumer

import (
	"fmt"
	"time"

	"github.com/harmony-one/go-raptorq/pkg/defaults"
	"github.com/harmony-one/go-raptorq/pkg/raptorq"
)

type ChunkRecord struct {
	Index        int
	CreateTime   int64
	Decoder      raptorq.Decoder
	ArrangePiece int
	PieceCount   int
	FinishChan   chan uint8
}

func NewChunkRecord(index int) (*ChunkRecord, error) {
	decode, err := defaults.NewDecoder(8, 8)
	if err != nil {
		return nil, err
	}

	finishChan := make(chan uint8)
	decode.AddReadyBlockChan(finishChan)

	return &ChunkRecord{
		Index:        index,
		CreateTime:   time.Now().Unix(),
		Decoder:      decode,
		FinishChan:   finishChan,
		ArrangePiece: 8,
	}, nil
}

func (c *ChunkRecord) IsReady() bool {
	select {
	case num := <-c.FinishChan:
		fmt.Println(num, "is ok", c.Decoder.SourceBlockSize(num))
		buf := make([]byte, 32, 32)
		n, err := c.Decoder.SourceBlock(num, buf)
		if err != nil {
			fmt.Println("wrong", err.Error())
			return true
		}
		fmt.Println(string(buf[:n]))
		return true
	default:
		return false
	}
}

func (c *ChunkRecord) Close() error {
	return c.Decoder.Close()
}

func (c *ChunkRecord) IsTimeout() bool {
	return false
}

func (c *ChunkRecord) AddPiece(sbn uint8, esi uint32, symbol []byte) {
	c.Decoder.Decode(sbn, esi, symbol)
}
