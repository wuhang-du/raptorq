package common

type TrackerInterface interface {
	RegisterRaqServer(RaqServerInterface)
	RegisterMicroServer(MicroInterface)
	GetRaqServer() RaqServerInterface
	GetMicroServer() MicroInterface
}