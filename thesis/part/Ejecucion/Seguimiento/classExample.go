package Example

import "errors"

type ClassName interface {
	GetVariable() string
	GetAnotherVariable() string
	SetVariable(newValue string)
}

type className struct {
	variable        string
	anotherVariable string
}

func NewClassName(variable string) (ClassName, error) {
	if variable == "invalid" {
		return nil, errors.New("error")
	}
	return &className{variable, "DEFAULT"}, nil
}

func (c *className) GetVariable() string {
	return c.variable
}

func (c *className) GetAnotherVariable() string {
	return c.anotherVariable
}

func (c *className) SetVariable(newValue string) {
	c.variable = newValue
}
