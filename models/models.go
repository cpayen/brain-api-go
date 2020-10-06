package models

import (
	"time"
)

type modelBase struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"cupdatedAt"`
}

type displayableModelBase struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// Folder is...
type Folder struct {
	modelBase
	displayableModelBase
}

// Note is...
type Note struct {
	modelBase
	displayableModelBase
	Content string `json:"content,omitempty"`
}

// Bookmark is...
type Bookmark struct {
	modelBase
	displayableModelBase
	URL string `json:"url"`
}
