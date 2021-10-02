package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/akhenakh/codingtestnr/tripletsvc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var (
	svcURI = flag.String("svcURI", "localhost:9200", "grpc URI")
	path   = flag.String("path", "", "a file to send for computation")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*svcURI,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name), //nolint:staticcheck
	)
	if err != nil {
		log.Fatal(err)
	}

	c := tripletsvc.NewTripletServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	b, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	triplets, err := c.Compute(ctx, &tripletsvc.ComputeRequest{Text: b})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", triplets)
}
