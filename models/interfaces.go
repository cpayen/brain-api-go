package models

import "time"

// Content is the base interface for all content types
type Content interface {
	SetCreationDate(d time.Time)
	SetUpdateDate(d time.Time)
}
