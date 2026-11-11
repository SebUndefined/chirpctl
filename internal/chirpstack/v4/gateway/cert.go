package gateway

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/gateway"
)

type cert struct {
	client Client
}

var _ gateway.Cert = (*cert)(nil)

func newCert(client Client) *cert {
	return &cert{client: client}
}

func (c *cert) GenerateClientCertificate(ctx context.Context, in gateway.GenerateGatewayClientCertificateParams) (gateway.GenerateGatewayClientCertificateResult, error) {
	//TODO implement me
	panic("implement me")
}
