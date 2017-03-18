package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/nickschuch/articles"
)

var (
	serverAddr = flag.String("api", "127.0.0.1:8080", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption

	// This is insecure, not for production!
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewArticlesClient(conn)

	articles := []*pb.Article{
		&pb.Article{
			Title: "This is a test",
			Body:  "Foo bar baz",
		},
	}

	// Seed the API with our data from above.
	for _, article := range articles {
		_, err := client.Create(context.Background(), &pb.CreateRequest{Article: article})
		if err != nil {
			log.Println("Failed to create article:", article.Title, err)
			continue
		}

		log.Println("Created article:", article.Title)
	}

	// Query and print out all the data that lives in the
	resp, err := client.List(context.Background(), &pb.ListRequest{})
	if err != nil {
		log.Fatalf("Failed to list articles:", err)
	}

	for _, article := range resp.Articles {
		fmt.Println(article.Title)
	}
}
