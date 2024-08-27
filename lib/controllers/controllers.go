package controllers

import (
	"context"
	"errors"
	"log"

	"github.com/inamandev/learning-ginkgo-with-grpc/lib/store"
	pb "github.com/inamandev/learning-ginkgo-with-grpc/students_proto"
)

type Server struct {
	pb.UnimplementedStudentsServer
}

func (*Server) Get(ctx context.Context, req *pb.StudentId) (*pb.Response, error) {
	// validate input
	// case 1 if id is zero or below
	if req.Id <= 0 {
		err := errors.New("invalid student id. Kindly provide valid id")
		log.Println(err.Error())
		return &pb.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}, err
	}
	student := &pb.StudentPayload{
		Id: req.Id,
	}
	err := store.DB.GetStudent(student)
	if err != nil {
		return &pb.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}, err

	}
	return &pb.Response{
		Success: true,
		Data:    student,
	}, nil
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
