/*
This package is reponsible for all the cli tasks for Node-runnner
like taking input, port addresses etc
printing errors ormessage
*/

package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jusxtdev/nodeRunner/internal/node"
)

// runs the cli app loop
func Run(nm *node.NodeManager){
	reader := bufio.NewReader(os.Stdin)
	// configuration prompt
	// - ask for number of nodes
	// - path of program entrypoint
	// - command to run the program
	var path string
	fmt.Printf("enter path of program to run :- ")
	path, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var command string
	fmt.Printf("enter command to execute program :- ")
	command, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var nodeCountstr string = "1"
	fmt.Printf("Enter number of nodes to create (default 1) :- ")
	nodeCountstr , err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	nodeCount, err := strconv.Atoi(nodeCountstr[0:1])
	if err != nil {
		log.Fatal(err)
	}

	// pass the inputs to nodeManager
	config := node.Configuration{
		NodeCount: nodeCount,
		ProgramPath: path,
		Command: command,
	}
	fmt.Printf("%#v\n", config)
	nm.SetConfig(config)
}