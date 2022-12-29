package Status

type Status string

const (
	Pending    Status = "PENDING"
	Running    Status = "RUNNING"
	Done       Status = "DONE"
	Successful Status = "SUCCESSFUL"
	failed     Status = "FAILED"
)

func FromString(mode string) (Status, error) {
	switch mode {
	case "PENDING":
		return Pending, nil
	case "RUNNING":
		return Running, nil
	case "DONE":
		return Done, nil
	case "SUCCESSFUL":
		return Successful, nil
	case "FAILED":
		return failed, nil
	}
	return "", NewUnknownError()
}
