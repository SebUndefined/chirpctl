package clientfactory

import "context"

type bearerCreds struct {
	token          string
	allowPlaintext bool
}

func (b bearerCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	if b.token == "" {
		return map[string]string{}, nil
	}

	return map[string]string{
		"authorization": "Bearer " + b.token,
	}, nil
}

func (b bearerCreds) RequireTransportSecurity() bool {
	return !b.allowPlaintext
}
