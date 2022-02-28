package ValueObject

import "errors"

type StepVo struct {
	Sentence string
}

func NewStepVo(sentence string) (StepVo, error) {
	self := StepVo{}
	if len(sentence) > 255 {
		return self, errors.New("step.sentence length must be less than 255")
	}
	self.Sentence = sentence
	return self, nil
}
