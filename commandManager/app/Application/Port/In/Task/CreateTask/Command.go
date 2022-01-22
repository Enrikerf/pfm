package CreateTask

type Command struct {
	Host          string
	Port          string
	Commands      []string
	Mode          string
	Status        string
	ExecutionMode string
}
