package node

type Configuration struct {
	NodeCount int
	ProgramPath string
	Command string
	Ports []int
	Nodes []Node
}

type NodeManager struct {
	Config Configuration
}

func NewManager()*NodeManager{
	return &NodeManager{}
}

func (nm *NodeManager) Initialize(config Configuration) {
	nm.Config.NodeCount = config.NodeCount
	nm.Config.ProgramPath = config.ProgramPath
	nm.Config.Command = config.Command
	nm.Config.Ports = config.Ports
	nm.initNodes()
}

func (nm *NodeManager) initNodes() {
	var ports []int = nm.Config.Ports
	var i int = 0 
	for i < nm.Config.NodeCount {
		var n Node
		n.NodeID = i+1
		n.Port = ports[i]
		n.Peers = sliceCpyExcludeindex(i, ports)
		nm.Config.Nodes = append(nm.Config.Nodes, n)
		i++
	}
}

/* - - - - - HELPERS  - - - - -  */
func sliceCpyExcludeindex(index int, original []int) []int {
	// creates a copy of a slice but excludes the element at the index provided
	// used to create slice of peers for i'th node hence excluding i'th port in the original ports slice 
	var cpy []int
	for i, elem := range original {
		if i == index {
			continue
		}
		cpy = append(cpy, elem)
	}
	return cpy
}