package core

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// MasterNode is the node instance
type MasterNode struct {
	api     *gin.Engine            // api server
	ln      net.Listener           // listener
	svr     *grpc.Server           // grpc server
	nodeSvr *NodeServiceGrpcServer // node service
}

func (n *MasterNode) Init() (err error) {
	// grpc server listener with port as 50051
	n.ln, err = net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	// grpc server
	n.svr = grpc.NewServer()

	// node service
	n.nodeSvr = GetNodeServiceGrpcServer()

	// register node service to grpc server
	RegisterNodeServiceServer(n.svr, n.nodeSvr)

	// api
	n.api = gin.Default()
	n.api.POST("/tasks", func(c *gin.Context) {
		var payload struct {
			Cmd string `json:"cmd"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		n.nodeSvr.CmdChannel <- payload.Cmd
		c.AbortWithStatus(http.StatusOK)
	})

	return nil
}

func (n *MasterNode) Start() {
	go n.svr.Serve(n.ln)

	_ = n.api.Run(":9092")

	n.svr.Stop()
}

var masterNode *MasterNode

func GetMasterNode() *MasterNode {
	if masterNode == nil {
		masterNode = &MasterNode{}

		if err := masterNode.Init(); err != nil {
			panic(err)
		}
	}

	return masterNode
}
