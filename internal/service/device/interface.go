package device

type API interface {
	Core
	Keys
	Activation
	Queue
	Metrics
	Link
}
