package Task

import "fmt"

type TaskModes int

const (
	Unary TaskModes = iota
	ServerStream
	ClientStream
	Bidirectional
)

func (taskModes TaskModes) String() string {
	switch taskModes {
	case Unary:
		return "Unary"
	case ServerStream:
		return "ServerStream"
	case ClientStream:
		return "ClientStream"
	case Bidirectional:
		return "Bidirectional"
	}
	return "unknown"
}

func GetTaskMode(mode string) (TaskModes, error) {
	switch mode {
	case "Unary":
		return Unary, nil
	case "ServerStream":
		return ServerStream, nil
	case "ClientStream":
		return ClientStream, nil
	case "Bidirectional":
		return Bidirectional, nil
	}
	return -1, fmt.Errorf("taskMode undefined")
}
