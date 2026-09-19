## Node-Runner
- A script that prompts user for 
  - program to run
  - the number of nodes to run that program
- Use case example
  - running 3 instances of a server program
- each instance of the program will have its own output log 
  - if there are 3 nodes printing something to os.Stdout
  - there will be 3 files - `node1.log` `node2.log`
  - i.e. each routine will manage getting the output from the program and writing to its logfile

### Some future ideas
- stop/restart all or few of the nodes

### Working
Runs `n` instances of a program in their own goroutine

#### Startup
- prompt for destination of the program
- also ask for the command to execute that program
  - Ex. - `go run ./other/main.go`
- ask for number of nodes to run for that program
- create a Node struct for each 
```go
type Node struct {
    NodeID int 
    Port int        // 8081 to run on localhost:8081
    Peers []int     // array of Port address of other Nodes
}
```

#### Running
- a goroutine for each Node is started 
- logging for each node is done in their separate log files
- the client (from the main program) is notified if any of the nodes has crashed or reported an error