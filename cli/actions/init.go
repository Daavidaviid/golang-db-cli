package actions

import (
	"blogCLI/models"
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func getScannedValue(scanner *bufio.Scanner, defaultValue string) string {
	scanner.Scan()
	value := scanner.Text()

	if value != "" {
		return value
	}

	return defaultValue
}

func Init(c *cli.Context) error {
	fmt.Println("INITIALIZING SKYBLOG")
	fmt.Println("--------------------")

	config := models.PostgresQLConfig{
		User:     "LePatron",
		Address:  "localhost",
		Password: "@&dtshcZShb66VX&77y!*s38n",
		Port:     "5432",
		Database: "skyblog_cli_lachot",
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Which PostgresQL user ? [%s]", config.User)
	config.User = getScannedValue(scanner, config.User)

	fmt.Printf("Which PostgresQL address ? [%s]", config.Address)
	config.Address = getScannedValue(scanner, config.Address)

	fmt.Printf("Which PostgresQL password ? [%s]", config.Password)
	config.Password = getScannedValue(scanner, config.Password)

	fmt.Printf("Which PostgresQL port ? [%s]", config.Port)
	config.Port = getScannedValue(scanner, config.Port)

	fmt.Printf("Which PostgresQL database ? [%s]", config.Database)
	config.Database = getScannedValue(scanner, config.Database)

	err := config.Save()

	fmt.Println("Configuration saved in .skyblogrc file !!!")

	return err
}
