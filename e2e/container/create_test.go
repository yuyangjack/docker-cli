package container

import (
	"fmt"
	"testing"

	"github.com/yuyangjack/dockercli/e2e/internal/fixtures"
	"github.com/yuyangjack/dockercli/internal/test/environment"
	"gotest.tools/icmd"
	"gotest.tools/skip"
)

func TestCreateWithContentTrust(t *testing.T) {
	skip.If(t, environment.RemoteDaemon())

	dir := fixtures.SetupConfigFile(t)
	defer dir.Remove()
	image := fixtures.CreateMaskedTrustedRemoteImage(t, registryPrefix, "trust-create", "latest")

	defer func() {
		icmd.RunCommand("docker", "image", "rm", image).Assert(t, icmd.Success)
	}()

	result := icmd.RunCmd(
		icmd.Command("docker", "create", image),
		fixtures.WithConfig(dir.Path()),
		fixtures.WithTrust,
		fixtures.WithNotary,
	)
	result.Assert(t, icmd.Expected{
		Err: fmt.Sprintf("Tagging %s@sha", image[:len(image)-7]),
	})
}

// FIXME(vdemeester) TestTrustedCreateFromBadTrustServer needs to be backport too (see https://github.com/moby/moby/pull/36515/files#diff-4b1e56bb77ac16f2ccf956fc24cf0a82L331)
