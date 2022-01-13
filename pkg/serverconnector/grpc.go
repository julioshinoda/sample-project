package serverconnector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	pb "github.com/julioshinoda/sample-project/pb/proto"
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
	conn, err := grpc.Dial(os.Getenv("SERVER_CLIENT"), transportOption)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}

func NewCustomerConnectorClient(cc *grpc.ClientConn) util.CustomerConnectorClient {
	return util.NewCustomerConnectorClient(cc)
}

//
func NewOrdersConnectorClient(cc *grpc.ClientConn) pb.OrdersConnectorClient {
	return pb.NewOrdersConnectorClient(cc)
}
func GetMetadata(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("token", os.Getenv("SERVER_TOKEN"), "clientuuid", os.Getenv("SERVER_CLIENT_UUID")),
	)

}
