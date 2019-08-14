package engine

import (
	clitypes "github.com/yuyangjack/docker-cli/types"
)

type extendedEngineInitOptions struct {
	clitypes.EngineInitOptions
	sockPath string
}
