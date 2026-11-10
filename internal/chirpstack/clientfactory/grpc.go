package clientfactory

import "context"

type bearerCreds struct {
	token    string
	insecure bool
}

func (b bearerCreds) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	if b.token == "" {
		return map[string]string{"user-agent": "chirpctl/1"}, nil
	}
	return map[string]string{
		"authorization": "Bearer " + b.token,
		"user-agent":    "chirpctl-1",
	}, nil
}
func (b bearerCreds) RequireTransportSecurity() bool { return !b.insecure }
