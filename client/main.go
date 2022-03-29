package main

import (
	"context"
	"flag"
	"log"
	"time"

	bookshelf "github.com/woodysee/small-grpc-go-service/bookshelf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Default Name"
)

var (
	addr = flag.String("addr", "localhost:8008", "the address to connect to")
	name = flag.String("name", defaultName, "Name of a default book")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bookshelf.NewBookHandlerClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBook(ctx, &bookshelf.GetBookRequest{Name: *name})
	if err != nil {
		log.Fatalf("Unable to return requested book from name: %v", err)
	}
	log.Printf("Book: %s", r.Book.Name)
}
