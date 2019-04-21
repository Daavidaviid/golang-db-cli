package article

import (
	"time"
)

type Article struct {
	tableName struct{}   `sql:"articles" `
	ID        string     `json:"id" faker:"uuid_hyphenated"`
	Name      string     `json:"firstName" faker:"first_name"`
	UserID    string     `json:"userId" faker:"-"`
	Content   string     `json:"lastName" faker:"paragraph"`
	DeletedAt *time.Time `json:"deletedAt" pg:",soft_delete" faker:"-"`
	CreatedAt time.Time  `json:"createdAt" faker:"-"`
	UpdatedAt time.Time  `json:"updatedAt" faker:"-"`
}
