package controllers

import (
	"context"
	"log"

	pb "github.com/inamandev/learning-ginkgo-with-grpc/students_proto"
)

type Server struct {
	pb.UnimplementedStudentsServer
}

func (*Server) Read(ctx context.Context, req *pb.StudentId) (*pb.Response, error) {
	log.Println(req.Email)
	return nil, nil
}
func (*Server) Create(ctx context.Context, req *pb.StudentPayload) (*pb.Response, error) {
	return nil, nil
}
func (*Server) Update(ctx context.Context, req *pb.StudentPayload) (*pb.Response, error) {
	return nil, nil
}
func (*Server) Delete(ctx context.Context, req *pb.StudentId) (*pb.Response, error) {
	return nil, nil
}
