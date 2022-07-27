package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

func doAuth() error {
	// migrations
	dbType := speed.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := speed.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := speed.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens cascade;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run Migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// Copy files over
	err = copyFileFromTemplate("templates/data/user.go.txt", speed.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", speed.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/remember_token.go.txt", speed.RootPath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth.go.txt", speed.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/remember.go.txt", speed.RootPath+"/middleware/remember.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", speed.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", speed.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/mailers/password-reset.plain.tmpl", speed.RootPath+"/mailers/password-reset.plain.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/mailers/password-reset.html.tmpl", speed.RootPath+"/mailers/password-reset.html.tmpl")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/forgot.jet", speed.RootPath+"/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/login.jet", speed.RootPath+"/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/views/reset-password.jet", speed.RootPath+"/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("  - users, tokens, remember_tokens migrations created and executed")
	color.Yellow("  - users and token model created")
	color.Yellow("  - auth middleware created")
	color.Yellow("")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and to add appropiate middleware to your routes!")

	return nil
}
