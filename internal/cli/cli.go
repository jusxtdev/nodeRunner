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
	"strings"
	"unicode"

	"github.com/jusxtdev/nodeRunner/internal/node"
)

// runs the cli app loop
func Run(nm *node.NodeManager){
	reader := bufio.NewReader(os.Stdin)

	// get the configuration for nodemanager
	// set the config in the manager
	config := InputConfig(reader)
	nm.Initialize(config)

	fmt.Printf("Configuration done\n CONFIG %v, NODES %v\n", nm.Config, nm.Config.Nodes)
}

func InputConfig(reader *bufio.Reader) node.Configuration {
	// configuration prompt
	// - ask for number of nodes
	// - path of program entrypoint
	// - command to run the program
	path := inputStringLine("Enter path of program to run :- ", reader)
	command := inputStringLine("Enter command to execute program :- ", reader)
	nodeCount := inputIntLine("Enter number of nodes to create (default 1) :- ", reader)
	// inut port addresses for 'NodeCount' nodes and set it in config 
	ports := InputPorts(reader, nodeCount)

	// pass the inputs to nodeManager
	config := node.Configuration{
		NodeCount: nodeCount,
		ProgramPath: path,
		Command: command,
		Ports: ports,
	}
	return config
}

func InputPorts(reader *bufio.Reader, nodeCount int) []int {
	var ports []int
	for i := range nodeCount {
		port := inputIntLine(fmt.Sprintf("Enter PORT to use for node %d :- ", i+1), reader)
		ports = append(ports, port)
	}
	return ports
}

/* - - - - - HELPERS  - - - - -  */
func inputStringLine(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimRightFunc(line, unicode.IsSpace)
}

func inputIntLine(prompt string, reader *bufio.Reader) int {
	line := inputStringLine(prompt, reader)

	value, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}

	return value
}