package engine

import (
	clitypes "github.com/yuyangjack/dockercli/types"
)

type extendedEngineInitOptions struct {
	clitypes.EngineInitOptions
	sockPath string
}
