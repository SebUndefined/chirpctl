package clients

import (
	"context"
	"crypto/tls"

	"github.com/sebundefined/chirpctl/internal/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func Dial(ctx context.Context, o options.Options) (*grpc.ClientConn, func() error, error) {
	var opts []grpc.DialOption

	if o.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tcfg := &tls.Config{
			MinVersion: tls.VersionTLS12,
			NextProtos: []string{"h2"},
		}
		if o.ServerName != "" {
			tcfg.ServerName = o.ServerName
		}
		// CACert: charge ton CA si fourni
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tcfg)))
	}

	// Authorization + UA
	unary := grpc.WithUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		callOpts ...grpc.CallOption,
	) error {
		md := metadata.New(nil)
		if o.Token != "" {
			md.Append("authorization", "Bearer "+o.Token)
		}
		md.Append("user-agent", "chipctl/1")
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, callOpts...)
	})
	opts = append(opts, unary)
	conn, err := grpc.NewClient(o.Addr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Close, nil
}
