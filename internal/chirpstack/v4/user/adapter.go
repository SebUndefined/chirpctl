package user

import (
	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/user"
	"google.golang.org/grpc"
)

type adapter struct {
	*core
}

var _ user.API = (*adapter)(nil)

func NewAdapter(conn *grpc.ClientConn) user.API {
	client := NewWithClient(apiv4.NewUserServiceClient(conn))
	return &adapter{
		core: newCore(client),
	}
}
