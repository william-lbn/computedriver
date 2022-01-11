package drvierchan

const (
	ClusterChanJoin   = "Join"
	ClusterChanDelete = "Delete"
	QuitChanInt       = 1
)

// ClusterChan join delete
var ClusterChan = make(chan string)

// QuitChan join delete
var QuitChan = make(chan int)
