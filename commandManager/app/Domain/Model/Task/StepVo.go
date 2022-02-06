package Task

type StepVo struct {
	Sentence string
}

func NewStepVo(sentence string) (StepVo, error) {
	self := StepVo{}
	self.Sentence = sentence
	return self, nil
}
