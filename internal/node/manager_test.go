package node

import "testing"

func TestInitializeBuildsNodesWithPeers(t *testing.T) {
	manager := NewManager()
	config := Configuration{
		NodeCount:   2,
		ProgramPath: "./server",
		Command:     "run",
		Ports:       []int{8081, 8082},
	}

	if err := manager.Initialize(config); err != nil {
		t.Fatalf("Initialize returned an error: %v", err)
	}

	if len(manager.Nodes) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(manager.Nodes))
	}
	if manager.Nodes[0].NodeID != 1 || manager.Nodes[0].Port != 8081 {
		t.Fatalf("unexpected first node: %#v", manager.Nodes[0])
	}
	if len(manager.Nodes[0].Peers) != 1 || manager.Nodes[0].Peers[0] != 8082 {
		t.Fatalf("unexpected first node peers: %#v", manager.Nodes[0].Peers)
	}
}

func TestInitializeRejectsMismatchedPorts(t *testing.T) {
	manager := NewManager()

	err := manager.Initialize(Configuration{NodeCount: 2, Ports: []int{8081}})
	if err == nil {
		t.Fatal("expected an error for mismatched ports")
	}
}
