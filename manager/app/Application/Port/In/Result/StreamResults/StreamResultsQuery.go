package StreamResults

import (
	"github.com/google/uuid"
	"time"
)

type Query struct {
	BatchUuid uuid.UUID
	LastDate  time.Time
}
