package system

import (
	"context"

	"github.com/yuyangjack/moby/api/types"
	"github.com/yuyangjack/moby/client"
)

type fakeClient struct {
	client.Client

	version       string
	serverVersion func(ctx context.Context) (types.Version, error)
}

func (cli *fakeClient) ServerVersion(ctx context.Context) (types.Version, error) {
	return cli.serverVersion(ctx)
}

func (cli *fakeClient) ClientVersion() string {
	return cli.version
}
