/*
This package is reponsible for all the cli tasks for Node-runnner
like taking input, port addresses etc
printing errors ormessage
*/

package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/jusxtdev/nodeRunner/internal/node"
)

// runs the cli app loop
func Run(nm *node.NodeManager) error {
	reader := bufio.NewReader(os.Stdin)

	config, err := InputConfig(reader)
	if err != nil {
		return err
	}

	if err := nm.Initialize(config); err != nil {
		return err
	}

	fmt.Println("Configuration done")
	return nil
}

func InputConfig(reader *bufio.Reader) (node.Configuration, error) {
	// configuration prompt
	// - ask for number of nodes
	// - path of program entrypoint
	// - command to run the program
	path, err := inputStringLine("Enter path of program to run :- ", reader)
	if err != nil {
		return node.Configuration{}, err
	}
	command, err := inputStringLine("Enter command to execute program :- ", reader)
	if err != nil {
		return node.Configuration{}, err
	}
	nodeCount, err := inputIntLine("Enter number of nodes to create (default 1) :- ", reader)
	if err != nil {
		return node.Configuration{}, err
	}
	ports, err := InputPorts(reader, nodeCount)
	if err != nil {
		return node.Configuration{}, err
	}

	return node.Configuration{
		NodeCount:   nodeCount,
		ProgramPath: path,
		Command:     command,
		Ports:       ports,
	}, nil
}

func InputPorts(reader *bufio.Reader, nodeCount int) ([]int, error) {
	var ports []int
	for i := range nodeCount {
		port, err := inputIntLine(fmt.Sprintf("Enter PORT to use for node %d :- ", i+1), reader)
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}
	return ports, nil
}

/* - - - - - HELPERS  - - - - -  */
func inputStringLine(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	line, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF && line != "" {
			return strings.TrimRightFunc(line, unicode.IsSpace), nil
		}
		return "", err
	}

	return strings.TrimRightFunc(line, unicode.IsSpace), nil
}

func inputIntLine(prompt string, reader *bufio.Reader) (int, error) {
	line, err := inputStringLine(prompt, reader)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(line)
	if err != nil {
		return 0, fmt.Errorf("invalid integer %q: %w", line, err)
	}

	return value, nil
}
