package Task

import "fmt"

type ExecutionMode int

const (
	Automatic ExecutionMode = iota
	Manual
)

func (executionMode ExecutionMode) String() string {
	switch executionMode {
	case Automatic:
		return "Automatic"
	case Manual:
		return "Manual"
	}
	return "unknown"
}

func GetExecutionMode(executionMode string) (ExecutionMode, error) {
	switch executionMode {
	case "Automatic":
		return Automatic, nil
	case "Manual":
		return Manual, nil
	}
	return -1, fmt.Errorf("undefined")
}
