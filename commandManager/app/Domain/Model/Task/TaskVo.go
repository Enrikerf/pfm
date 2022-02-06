package Task

type TaskVo struct {
	Host          string
	Port          string
	Commands      []Step
	Mode          Modes
	Status        TaskStatus
	ExecutionMode ExecutionMode
}

func NewTaskVo(host string, port string, commands []Step, mode string, executionMode string) (TaskVo, error) {
	taskVo := TaskVo{}
	taskVo.Host = host
	taskVo.Port = port
	taskVo.Commands = commands
	taskMode, err := GetTaskMode(mode)
	if err != nil {
		return TaskVo{}, err
	}
	taskExecutionMode, err := GetExecutionMode(executionMode)
	if err != nil {
		return TaskVo{}, err
	}
	taskVo.ExecutionMode = taskExecutionMode
	taskVo.Mode = taskMode
	taskVo.Status = Pending
	return taskVo, nil
}
