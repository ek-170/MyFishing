package models

import (
	"time"
)

type Users struct {
	id        int
	ulid      string
	email     string
	createdAt time.Time
}
