package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup(arg1, arg2 string) {
	if arg1 != "new" && arg1 != "version" && arg1 != "help" {
		err := godotenv.Load()
		if err != nil {
			exitGracefully(err)
		}

		path, err := os.Getwd()
		if err != nil {
			exitGracefully(err)
		}
		speed.RootPath = path
		speed.DB.DataType = os.Getenv("DATABASE_TYPE")
	}
}

func getDSN() string {
	dbType := speed.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn
	}
	return "mysql://" + speed.BuildDSN()
}

func showHelp() {
	color.Yellow(`Available Commands:

	help                     - Show help commands.
	version                  - Show Speedlight version.
	migrate                  - Run all migrations that had not been run previously.
	migrate down             - Reverses the most recent migration.
	migrate reset            - Run down all migration in reverse order, and the all up migrations.
	make migration <name>    - Creates Up and Down Migrations at migration folder.
	make auth                - Creates and run migrations for authentication tables and creates models and middleware.
	make handler <name>      - Creates a stub handler in the handler directory.
	make model <name>        - Creates a new data in model directory.
	make session             - Creates a table in the database as a session store.
	make mail <name>         - Creates two starter mail templates in the mail directory.


	`)
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	// check the error before do anything else
	if err != nil {
		exitGracefully(err)
	}

	// check if current file is directory
	if fi.IsDir() {
		return nil
	}

	// only check go files
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	// we have a matching file
	if matched {
		// read file contents
		read, err := os.ReadFile(path)
		if err != nil {
			exitGracefully(err)
		}
		newContents := strings.Replace(string(read), "App/", appURL+"/", -1)

		// write the changed files
		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}

func updateSource() {
	// walk entire project folder including subfolders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		exitGracefully(err)
	}
}
