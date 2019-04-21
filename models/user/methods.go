package user

import "fmt"

func (u User) String() string {
	return fmt.Sprintf("#%s <%s> - %s %s (since %s)", u.ID, u.Email, u.FirstName, u.LastName, u.CreatedAt.Format("Monday Jan _2 at 15:04"))
}
