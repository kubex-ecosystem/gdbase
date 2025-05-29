package interfaces

type ISignalManager[T chan string] interface {
	ListenForSignals() error
	StopListening()
}
