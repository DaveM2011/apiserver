package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/davem2011/apiserver/modules/visitorip"
	_ "github.com/davem2011/apiserver/modules/bookingapi"
)

func main() {
	caddycmd.Main()
}
