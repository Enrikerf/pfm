package ExecutionMode

type ExecutionMode string

const (
	Manual    ExecutionMode = "MANUAL"
	Automatic ExecutionMode = "AUTOMATIC"
)

func FromString(mode string) (ExecutionMode, error) {
	switch mode {
	case "MANUAL":
		return Manual, nil
	case "AUTOMATIC":
		return Automatic, nil
	}
	return "", NewUnknownError()
}
