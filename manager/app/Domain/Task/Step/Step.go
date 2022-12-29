package Step

import (
	"github.com/google/uuid"
)

type Step interface {
	GetUuid() uuid.UUID
	GetUuidString() string
	GetSentence() string
}

type step struct {
	uuid     uuid.UUID
	sentence string
}

func (s step) GetUuid() uuid.UUID {
	return s.uuid
}

func (s step) GetUuidString() string {
	return s.uuid.String()
}

func (s step) GetSentence() string {
	return s.sentence
}

func New(stepVo Vo) Step {
	self := &step{}
	self.uuid = uuid.New()
	self.sentence = stepVo.GetSentence()
	return self
}
