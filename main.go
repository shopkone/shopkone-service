package main

import (
	_ "shopkone-service/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shopkone-service/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
