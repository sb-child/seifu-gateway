package main

import (
	_ "seifu-gw/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"seifu-gw/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
