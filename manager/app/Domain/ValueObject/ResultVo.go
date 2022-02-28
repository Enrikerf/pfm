package ValueObject

import (
	"time"
)

type ResultVo struct {
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewResultVo(content string) ResultVo {
	resultVo := ResultVo{}
	resultVo.Content = content
	resultVo.CreatedAt = time.Now()
	resultVo.UpdatedAt = time.Now()
	return resultVo
}
