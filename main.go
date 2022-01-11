package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	pb "github.com/julioshinoda/sample-project/pb/proto"
	"github.com/julioshinoda/sample-project/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	// add key-value pairs of metadata to context
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("token", "storer@2020!", "clientuuid", "fdfd497828effdfda2b049a27849a2d4"),
	)

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
	defer conn.Close()
	client := util.NewCustomerConnectorClient(conn)
	email := &util.Email{Email: "gehrmann.j@nekom.com"}
	resp, err := client.GetCustomerByEmail(ctx, email)
	if err != nil {
		fmt.Println("errror =/", err)
		return

	}
	fmt.Printf("resp:%+v <|", resp)
	orderCl := pb.NewOrdersConnectorClient(conn)
	fmt.Println("orderCl =/", orderCl)

	customerID := &pb.CustomerId{Customerid: int32(resp.GetCustomers()[0].Customerid)}
	orderResponse, orderErr := orderCl.GetOrdersForCustomer(ctx, customerID)
	if orderErr != nil {
		fmt.Println("errror =/", orderErr.Error())

	}
	fmt.Printf("orderResponse:%v <|", orderResponse)
}
