package grpc

import (
	"io"
	"log"

	protoport "justCode/Lecture18/proto/gw"
)

type Port struct {
	server protoport.PortServiceServer
	protoport.UnimplementedPortServiceServer
}

func (p Port) Create(stream protoport.PortService_CreateServer) error {
	var total int32

	for {
		port, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protoport.CreateResponse{
				Total: total,
			})
		}
		if err != nil {
			return err
		}

		total++
		log.Printf("%+v\n", port)
	}
}
