package gateway

import "github.com/sebundefined/chirpctl/internal/service/shared"

type API interface {
	Core
	Cert
	DutyCycle
	Relay
	shared.Metrics[GetGatewayMetricsParams, GetGatewayMetricsResult]
}
