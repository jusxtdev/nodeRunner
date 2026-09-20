package node

import "fmt"

type Configuration struct {
	NodeCount   int
	ProgramPath string
	Command     string
	Ports       []int
}

type NodeManager struct {
	Config Configuration
	Nodes  []Node
}

func NewManager() *NodeManager {
	return &NodeManager{}
}

func (nm *NodeManager) Initialize(config Configuration) error {
	if config.NodeCount <= 0 {
		return fmt.Errorf("node count must be greater than zero")
	}
	if len(config.Ports) != config.NodeCount {
		return fmt.Errorf("expected %d ports, got %d", config.NodeCount, len(config.Ports))
	}

	nm.Config = config
	nm.Nodes = buildNodes(config.Ports)
	return nil
}

func buildNodes(ports []int) []Node {
	nodes := make([]Node, 0, len(ports))
	for i, port := range ports {
		nodes = append(nodes, Node{
			NodeID: i + 1,
			Port:   port,
			Peers:  sliceCopyExcludeIndex(i, ports),
		})
	}
	return nodes
}

// func (nm *NodeManager)StartAllNodes(){
// 	for _, n := range nm.Nodes {

// 	}
// }

/* HELPERS */
func sliceCopyExcludeIndex(index int, original []int) []int {
	// creates a copy of a slice but excludes the element at the index provided
	// used to create slice of peers for i'th node hence excluding i'th port in the original ports slice
	copyOfSlice := make([]int, 0, len(original)-1)
	for i, elem := range original {
		if i == index {
			continue
		}
		copyOfSlice = append(copyOfSlice, elem)
	}
	return copyOfSlice
}
