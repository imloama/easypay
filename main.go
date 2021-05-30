package main

import (
	_ "github.com/imloama/easypay/boot"
	_ "github.com/imloama/easypay/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
