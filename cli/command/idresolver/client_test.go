package idresolver

import (
	"context"

	"github.com/yuyangjack/moby/api/types"
	"github.com/yuyangjack/moby/api/types/swarm"
	"github.com/yuyangjack/moby/client"
)

type fakeClient struct {
	client.Client
	nodeInspectFunc    func(string) (swarm.Node, []byte, error)
	serviceInspectFunc func(string) (swarm.Service, []byte, error)
}

func (cli *fakeClient) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error) {
	if cli.nodeInspectFunc != nil {
		return cli.nodeInspectFunc(nodeID)
	}
	return swarm.Node{}, []byte{}, nil
}

func (cli *fakeClient) ServiceInspectWithRaw(ctx context.Context, serviceID string, options types.ServiceInspectOptions) (swarm.Service, []byte, error) {
	if cli.serviceInspectFunc != nil {
		return cli.serviceInspectFunc(serviceID)
	}
	return swarm.Service{}, []byte{}, nil
}
