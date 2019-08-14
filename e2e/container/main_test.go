package container

import (
	"fmt"
	"os"
	"testing"

	"github.com/yuyangjack/docker-cli/internal/test/environment"
)

func TestMain(m *testing.M) {
	if err := environment.Setup(); err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	os.Exit(m.Run())
}
