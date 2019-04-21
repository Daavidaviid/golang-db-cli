package user

import (
	"blogCLI/database"
	"blogCLI/env"
	"errors"
	"fmt"

	"github.com/gobuffalo/uuid"
	"github.com/urfave/cli"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(env.AppPasswordSalt+password), 14)
	return string(bytes), err
}

func Create(c *cli.Context) error {
	userUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	userUUIDString := userUUID.String()

	userPhone := c.String("phone")
	hashedPassword, err := hashPassword(c.String("password"))
	if err != nil {
		return err
	}

	newUser := &User{
		ID:             userUUIDString,
		FirstName:      c.String("firstName"),
		LastName:       c.String("lastName"),
		Email:          c.String("email"),
		HashedPassword: hashedPassword,
		Phone:          &userPhone,
	}

	err = database.DB.Insert(newUser)
	if err != nil {
		return err
	}

	fmt.Printf("Created user : %s", newUser)

	return nil
}

func Read(c *cli.Context) error {
	userUUIDString := c.String("id")
	if userUUIDString == "" {
		return errors.New("MISSING_ID")
	}

	readUser := &User{
		ID: userUUIDString,
	}

	err := database.DB.Select(readUser)
	if err != nil {
		return err
	}

	fmt.Printf("Read user : %s\n", readUser)

	return nil
}

func Update(c *cli.Context) error {

	userUUIDString := c.String("id")
	if userUUIDString == "" {
		return errors.New("MISSING_ID")
	}

	updatedUser := &User{
		ID:        userUUIDString,
		FirstName: c.String("firstName"),
	}

	_, err := database.DB.Model(updatedUser).WherePK().UpdateNotNull()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATE - USER #%s SUCCESS\n", updatedUser.ID)
	return nil
}

func Delete(c *cli.Context) error {
	userUUIDString := c.String("id")
	if userUUIDString == "" {
		return errors.New("MISSING_ID")
	}

	deletedUser := &User{
		ID: userUUIDString,
	}

	err := database.DB.Delete(deletedUser)
	if err != nil {
		return err
	}

	fmt.Printf("DELETE - USER #%s SUCCESS\n", deletedUser.ID)

	return nil
}

func List(c *cli.Context) error {
	var users []User

	err := database.DB.Model(&users).Order("created_at ASC").Select()

	fmt.Printf("Users :\n")
	for _, user := range users {
		fmt.Println(user)
	}

	return err
}
