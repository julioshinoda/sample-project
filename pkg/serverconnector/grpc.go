package serverconnector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"

	"github.com/julioshinoda/sample-project/util"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func NewConnector() *grpc.ClientConn {
	roots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal("failed to get system certificate pool")
	}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            roots,
	}
	transportOption := grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
	conn, err := grpc.Dial("stage.nekom.com:443", transportOption)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}

func NewCustomerConnectorClient(cc *grpc.ClientConn) util.CustomerConnectorClient {
	return util.NewCustomerConnectorClient(cc)
}

func GetMetadata(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("token", "storer@2020!", "clientuuid", "fdfd497828effdfda2b049a27849a2d4"),
	)

}
