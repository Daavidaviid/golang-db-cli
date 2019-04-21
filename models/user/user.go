package user

import (
	"time"
)

type User struct {
	tableName      struct{}   `sql:"users" `
	ID             string     `json:"id" faker:"uuid_hyphenated"`
	FirstName      string     `json:"firstName" faker:"first_name"`
	LastName       string     `json:"lastName" faker:"last_name"`
	HashedPassword string     `json:"hashedPassword" faker:"uuid_digit"`
	PhotoURL       *string    `json:"photoUrl" faker:"url"`
	Email          string     `json:"email" faker:"email"`
	Phone          *string    `json:"phone" faker:"e_164_phone_number"`
	DeletedAt      *time.Time `json:"deletedAt" pg:",soft_delete" faker:"-"`
	CreatedAt      time.Time  `json:"createdAt" faker:"-"`
	UpdatedAt      time.Time  `json:"updatedAt" faker:"-"`
}
