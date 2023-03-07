package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	nanoid "github.com/matoous/go-nanoid/v2"
	"google.golang.org/grpc"
	f "tuwien.ac.at/fusevin/fuse"
	"tuwien.ac.at/fusevin/vin"
	pb "tuwien.ac.at/fusevin/vin"
)

type GRPCServer struct {
	Port int
	Fuseserver *f.FuseServer
	WaitGroup *sync.WaitGroup

	server *server
}

func (gsrv *GRPCServer) Start() error {
	defer gsrv.WaitGroup.Done()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gsrv.Port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	gsrv.server = &server{fuseServer: gsrv.Fuseserver, channelMap: make(map[string]chan int)}
	pb.RegisterFuseVinServer(s, gsrv.server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (gsrv *GRPCServer) Notify(id string) error {
	gsrv.server.channelMap[id] <- 1
	return nil
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedFuseVinServer

	fuseServer *f.FuseServer
	channelMap map[string]chan int
}

// SayHello implements helloworld.GreeterServer
func (s *server) CreatePuppet(ctx context.Context, in *pb.Empty) (*pb.CreatePuppetReply, error) {
	id, err := nanoid.New()
	log.Printf("Creating special file " + id)
	s.fuseServer.CreateFile(id)

	return &pb.CreatePuppetReply{Id: id, VinFilename: s.fuseServer.Mountpoint + "/" + id}, err
}

func (s *server) StartStdinNotify(request *pb.StartStdinNotifyRequest, stream pb.FuseVin_StartStdinNotifyServer) error {
	//add channel to map
	s.channelMap[request.GetId()] = make(chan int)

	for {
		//blocking read for specific id
	  <-s.channelMap[request.GetId()]
	
		if err := stream.Send(&pb.StdinNotify{}); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) SupplyStdinContent(ctx context.Context, in *pb.StdinContent) (*vin.Empty, error) {
	s.fuseServer.SupplyStdin(in.Id, in.Payload)
	return &pb.Empty{}, nil
}

func (s *server) DestroyPuppet(ctx context.Context, in *pb.DestroyPuppetRequest) (*vin.Empty, error) {
	s.fuseServer.DestroyFile(in.Id)
	return &pb.Empty{}, nil
}
