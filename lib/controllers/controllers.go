package controllers

import (
	"context"
	pb "ginkgo-grpc-leanring-server-for-interview/students_proto"
	"log"
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
