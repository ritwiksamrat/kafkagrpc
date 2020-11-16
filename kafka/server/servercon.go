package main

import (
	"context"
	"log"
	"net"
	blogpb "path/kafka/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type blogItem struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type BlogServiceServer struct {
}

func main() {

	fmt.Println("Hi!")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterBlogServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {

	db, err := db.getconn()
	if err != nil {
		log.Println("Something Went Wrong")
	}

	userid := req.GetId()
	username := req.GetUsername()
	name := req.GetName()

	result, err := db.Query("insert into new_table values(?,?,?,?)", userid, username, name)
	if err != nil {
		panic(err.Error())
	}
	response := &blogpb.CreateBlogRes{Result: result}
	defer db.Close()
	return response, nil
}
