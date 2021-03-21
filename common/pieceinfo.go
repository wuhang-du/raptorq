package common

type PieceInfo struct {
	chid int64
	piece []byte
} 

func (p *PieceInfo) GetChunkId() int64 {
	return p.chid
}

func (p *PieceInfo) GetPiece() []byte {
	return p.piece
}