package ValueObject

import "fmt"

type TaskStatus int

const (
	Pending TaskStatus = iota
	Running
	Done
)

func (taskStatus TaskStatus) String() string {
	switch taskStatus {
	case Pending:
		return "Pending"
	case Running:
		return "Running"
	case Done:
		return "Done"
	}
	return "unknown"
}

func GetStatus(status string) (TaskStatus, error) {
	switch status {
	case "Pending":
		return Pending, nil
	case "Running":
		return Running, nil
	case "Done":
		return Done, nil
	}
	return -1, fmt.Errorf("undefined")
}
