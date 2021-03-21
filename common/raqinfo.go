package common

type RaqServerInterface interface {
	GetRaqInfo() (uint64, uint32)
	RegisterPiece(string, int64, int64, chan []PieceInfo) error
	MissData(string, int64, int64) ([]PieceInfo, error)
}
