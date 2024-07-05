package main

import (
	"github.com/joshua468/myapp/api/handler"
	"github.com/vercel/go-bridge/go/bridge"
)

func main() {
	bridge.Start(handler.Handler)
}
