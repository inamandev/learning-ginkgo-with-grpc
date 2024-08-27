package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime/debug"
	"time"

	"github.com/inamandev/learning-ginkgo-with-grpc/lib/controllers"
	"github.com/inamandev/learning-ginkgo-with-grpc/lib/store"
	pb "github.com/inamandev/learning-ginkgo-with-grpc/students_proto"

	"google.golang.org/grpc"
)

var (
	grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcHandlerInterceptor()))
	port       = flag.Int("port", 4567, "The server port")
)

func main() {
	flag.Parse()
	bootstart()
}

func bootstart() {
	defer restart()

	err := store.Connect()
	if err != nil {
		log.Println("unlable to connect to database: ", err)
		panic(err)
	}
	initServer()
}

func initServer() {
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Panicln("unable to setup a listen socket: ", err)
	}
	pb.RegisterStudentsServer(grpcServer, &controllers.Server{})
	log.Println("server is running on port: ", *port)
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Panicln("unable to start grpc server: ", err)
	}
}

func grpcHandlerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		panicked := true
		defer grpcRecovery(panicked)
		res, err := handler(ctx, req)
		panicked = false
		return res, err
	}
}

func grpcRecovery(panicked bool) {
	if panicked {
		if ok := recover(); ok != nil {
			log.Printf("got panic: %s", string(debug.Stack()))
		}
	}
}

func restart() {
	e := recover()
	log.Println(e.(error))
	if err, ok := e.(error); ok {
		log.Println("restarting the app due to: ", err)
		time.Sleep(time.Second * 10)
		bootstart()
	}
}
