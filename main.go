package main

import (
	"net/http"

	"github.com/joshua468/myapp/api/handler"
	"github.com/vercel/go-bridge/go/bridge"
)

func main() {

	bridge.Start(http.HandlerFunc(handler.Handler))
}
