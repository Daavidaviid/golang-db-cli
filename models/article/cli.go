package article

import (
	"blogCLI/database"
	"blogCLI/models/user"
	"errors"
	"fmt"

	"github.com/gobuffalo/uuid"
	"github.com/urfave/cli"
)

func getUserID(c *cli.Context) (string, error) {
	var userID string

	if c.String("userEmail") != "" {
		wantedUser := &user.User{}

		err := database.DB.Model(wantedUser).Where("email = ?", c.String("userEmail")).Select()
		if err != nil {
			return "", err
		}
		userID = wantedUser.ID
	} else if c.String("userId") != "" {
		userID = c.String("userId")
	}

	return userID, nil
}

func Create(c *cli.Context) error {
	var userID string

	articleUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	userID, err = getUserID(c)
	if err != nil {
		return err
	}

	articleUUIDString := articleUUID.String()

	newArticle := &Article{
		ID:      articleUUIDString,
		Name:    c.String("name"),
		UserID:  userID,
		Content: c.String("content"),
	}

	err = database.DB.Insert(newArticle)
	if err != nil {
		return err
	}

	fmt.Printf("CREATE - ARTICLE #%s SUCCESS\n", articleUUIDString)

	return nil
}

func Read(c *cli.Context) error {
	userUUIDString := c.String("id")
	if userUUIDString == "" {
		return errors.New("MISSING_ARTICLE_ID")
	}

	readArticle := &Article{
		ID: userUUIDString,
	}

	err := database.DB.Select(readArticle)
	if err != nil {
		return err
	}

	fmt.Printf("Read article : %s\n", readArticle)

	return nil
}

func Update(c *cli.Context) error {
	var userID string

	articleUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	userID, err = getUserID(c)
	if err != nil {
		return err
	}

	articleUUIDString := articleUUID.String()

	newArticle := &Article{
		ID:      c.String("id"),
		Name:    c.String("name"),
		UserID:  userID,
		Content: c.String("content"),
	}

	err = database.DB.Update(newArticle)
	if err != nil {
		return err
	}

	fmt.Printf("UPDATE - ARTICLE #%s SUCCESS\n", articleUUIDString)

	return nil
}

func Delete(c *cli.Context) error {
	userUUIDString := c.String("id")
	if userUUIDString == "" {
		return errors.New("MISSING_ID")
	}

	deletedArticle := &Article{
		ID: userUUIDString,
	}

	err := database.DB.Delete(deletedArticle)
	if err != nil {
		return err
	}

	fmt.Printf("DELETE - ARTICLE #%s SUCCESS\n", deletedArticle.ID)

	return nil
}

func List(c *cli.Context) error {
	var articles []Article

	err := database.DB.Model(&articles).Order("created_at ASC").Select()

	fmt.Printf("Articles :\n")
	if len(articles) > 0 {
		for _, article := range articles {
			fmt.Println(article)
		}
	} else {
		fmt.Println("None")
	}

	return err
}
