package deviceprofile

import (
	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
	"google.golang.org/grpc"
)

type adapter struct {
	*core
}

var _ deviceprofile.API = (*adapter)(nil)

func NewAdapter(conn *grpc.ClientConn) deviceprofile.API {
	client := NewWithClient(apiv4.NewDeviceProfileServiceClient(conn))
	return &adapter{
		core: newCore(client),
	}
}
