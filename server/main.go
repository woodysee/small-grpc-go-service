package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	bookshelf "github.com/woodysee/small-grpc-go-service/bookshelf"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8008, "The server port")
)

// Server is used to implement bookshelf.BookServiceServer.
type server struct {
	bookshelf.UnimplementedBookHandlerServer
}

func (s *server) GetBook(ctx context.Context, in *bookshelf.GetBookRequest) (*bookshelf.GetBookResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &bookshelf.GetBookResponse{
		Book: &bookshelf.Book{
			Name: in.Name,
			Isbn: 1231,
			Author: &bookshelf.Author{
				Name:              "You",
				YearOfPublication: "2022",
			},
		},
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bookshelf.RegisterBookHandlerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
