package main

import (
	"github.com/jusxtdev/nodeRunner/internal/cli"
	"github.com/jusxtdev/nodeRunner/internal/node"
)

func main(){
	// initialize node manager
	nm := node.NewManager()
	// start cli application
	cli.Run(nm)
	// wait for goroutines
}