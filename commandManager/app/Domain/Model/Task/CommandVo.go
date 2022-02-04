package Task

type CommandVo struct {
	Name string
}

func NewCommandVo(name string) (CommandVo, error) {
	commandVo := CommandVo{}
	commandVo.Name = name
	return commandVo, nil
}
