package main

import (
	"github.com/MackanThe1/d7024e-tutorial/internal/cli"
	"github.com/MackanThe1/d7024e-tutorial/pkg/build"
)

var (
	BuildVersion string = ""
	BuildTime    string = ""
)

func main() {
	build.BuildVersion = BuildVersion
	build.BuildTime = BuildTime
	cli.Execute()
}
