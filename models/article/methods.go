package article

import (
	"blogCLI/database"
	"blogCLI/models/user"
	"fmt"
	"log"
)

func (a Article) String() string {
	author := &user.User{
		ID: a.UserID,
	}

	err := database.DB.Select(author)
	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("#%s - %s : %s (%s)\n\tBy : %s", a.ID, a.Name, a.Content, a.CreatedAt.Format("Monday Jan _2 at 15:04"), author)
}
