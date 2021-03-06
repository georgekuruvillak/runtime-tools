package main

import (
	"os"
	"runtime"

	"github.com/opencontainers/runtime-tools/validation/util"
)

func main() {
	if "linux" != runtime.GOOS {
		util.Skip("linux-specific process.capabilities test", map[string]string{"OS": runtime.GOOS})
		os.Exit(0)
	}

	g := util.GetDefaultGenerator()
	g.SetupPrivileged(true)
	err := util.RuntimeInsideValidate(g, nil)
	if err != nil {
		util.Fatal(err)
	}
}
