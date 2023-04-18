package core

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
)

type WorkerNode struct {
	conn *grpc.ClientConn
	c    NodeServiceClient
}

func (n *WorkerNode) Init() (err error) {
	n.conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}

	n.c = NewNodeServiceClient(n.conn)

	return nil
}

func (n *WorkerNode) Start() {
	fmt.Println("worker node started")
	_, _ = n.c.ReportStatus(context.Background(), &Request{})

	stream, _ := n.c.AssignTask(context.Background(), &Request{})
	for {
		res, err := stream.Recv()
		if err != nil {
			return
		}

		fmt.Println("received command: ", res.Data)

		parts := strings.Split(res.Data, " ")
		if err := exec.Command(parts[0], parts[1:]...).Run(); err != nil {
			fmt.Println(err)
		}
	}
}

var workerNode *WorkerNode

func GetWorkerNode() *WorkerNode {
	if workerNode == nil {
		workerNode = &WorkerNode{}

		if err := workerNode.Init(); err != nil {
			panic(err)
		}
	}

	return workerNode
}
