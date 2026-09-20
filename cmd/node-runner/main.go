package main

import (
	"log"

	"github.com/jusxtdev/nodeRunner/internal/cli"
	"github.com/jusxtdev/nodeRunner/internal/node"
)

func main() {
	nm := node.NewManager()
	if err := cli.Run(nm); err != nil {
		log.Fatal(err)
	}
}
