package gateway

import (
	"context"
)

type Cert interface {
	GenerateClientCertificate(ctx context.Context, in GenerateGatewayClientCertificateParams) (GenerateGatewayClientCertificateResult, error)
}

type GenerateGatewayClientCertificateParams struct{}

type GenerateGatewayClientCertificateResult struct {
}
