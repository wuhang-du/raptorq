package common

type MicroInterface interface {
	MicroRegisterPiece(string,int64,
		int64,chan []PieceInfo) error
}

