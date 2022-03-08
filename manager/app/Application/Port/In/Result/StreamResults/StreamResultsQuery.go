package StreamResults

import (
	"github.com/google/uuid"
)

type Query struct {
	BatchUuid uuid.UUID
	LastId    uint
}
