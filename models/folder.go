package models

import "time"

// Folder is...
type Folder struct {
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
}

// NewFolder is...
// func NewFolder(title string) *Folder {
// 	f := new(Folder)
// 	f.Type = "folder"
// 	f.Title = title
// 	return f
// }

// SetCreationDate is...
func (c *Folder) SetCreationDate(d time.Time) {
	c.CreatedAt = d
}

// SetUpdateDate is...
func (c *Folder) SetUpdateDate(d time.Time) {
	c.UpdatedAt = d
}
