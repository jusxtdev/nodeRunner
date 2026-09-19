package node

type Configuration struct {
	NodeCount int
	ProgramPath string
	Command string
}

type NodeManager struct {
	Config Configuration
}

func NewManager()*NodeManager{
	return &NodeManager{}
}

func (nm *NodeManager) SetConfig(config Configuration) {
	nm.Config.NodeCount = config.NodeCount
	nm.Config.ProgramPath = config.ProgramPath
	nm.Config.Command = config.Command
}