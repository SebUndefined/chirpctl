package clientfactory

type APIVersion string

const (
	APIVersionV4 APIVersion = "4"
	// APIVersionV3 APIVersion = "3" // TODO: Add v3 adapter.
)

type Transport string

const (
	TransportGRPC Transport = "grpc"
	// TransportREST Transport = "rest" // TODO: When supported.
)
