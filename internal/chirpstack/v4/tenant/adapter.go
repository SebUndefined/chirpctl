package tenant

import (
	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
	"google.golang.org/grpc"
)

type adapter struct {
	*core
}

var _ tenant.API = (*adapter)(nil)

func NewAdapter(conn *grpc.ClientConn) tenant.API {
	client := NewWithClient(apiv4.NewTenantServiceClient(conn))
	return &adapter{
		core: newCore(client),
	}
}
