package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/zemirco/uid"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	pb "github.com/nickschuch/articles"
)

var port = flag.Int("port", 8080, "The server port")

type articlesServer struct {
	articles map[string]*pb.Article
}

func (a *articlesServer) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	// Generate a new hash for this new peice of content.
	u := uid.New(20)

	// Add it to this APIs in memory storage.
	a.articles[u] = in.Article

	// Pass back the hash, so the user has an ID to what they created.
	return &pb.CreateResponse{Id: u}, nil
}

func (a *articlesServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	// Do we have this item of content? If we do we should return it.
	if val, ok := a.articles[in.Id]; ok {
		return &pb.GetResponse{Article: val}, nil
	}

	return nil, fmt.Errorf("Cannot find content:", in.Id)
}

func (a *articlesServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// Loop over all the provided IPs, and remove the article if it exists.
	for _, i := range in.Ids {
		if _, ok := a.articles[i]; ok {
			delete(a.articles, i)
		}
	}

	return nil, nil
}

func (a *articlesServer) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	resp := new(pb.ListResponse)

	// Rebuild the list from a map to a slice.
	//   eg. map[string]*pb.Article to []*pb.Article
	for _, a := range a.articles {
		resp.Articles = append(resp.Articles, a)
	}

	return resp, nil
}

func main() {
	flag.Parse()

	var opts []grpc.ServerOption

	log.Println("Starting Articles API server")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new server.
	s := new(articlesServer)
	s.articles = make(map[string]*pb.Article)

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArticlesServer(grpcServer, s)
	grpcServer.Serve(listen)
}
