package main

import (
	"fmt"
	"github.com/Pandademic/luve"
	"github.com/go-git/go-git/v5"
	"github.com/pandademic/raspberry"
	"os"
)

var (
	UserHomeDir string
	UserTsdmDir string
)

func setupDirs() {
	doesDirExist, err := os.Stat(UserTsdmDir)
	luve.Luve(doesDirExist)
	if os.IsNotExist(err) {
		if err := os.Mkdir(UserTsdmDir, os.ModePerm); err != nil {
			fmt.Println("error creating directorys")
			luve.Luve(err)
			os.Exit(1)
		} else {
			fmt.Println("Success!")
		}
	} else {
		fmt.Println("Directory already exists!")
		os.Exit(1)
	}
}
func getRepo() {
	cloneDir := UserTsdmDir + string(os.PathSeparator) + "dotfile-repo"
	fmt.Println("getting a repository from ", raspberry.Args[0], " to: ", cloneDir)
	_, err := git.PlainClone(cloneDir, false, &git.CloneOptions{
		URL:      raspberry.Args[0],
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	help := `
           TSDM HELP 
         --------------

         the simple dotfile manager

         AVAILABLE COMMANDS:

         - help <- show this help message

         - -h <- show this help message

         - version <- show the version number of tsdm

         - -v <- show the version number of tsdm
        
         - setup <- setup the tsdm directory structure

         - get <- get a dotfile repo to store in your tsdm directory.Note that this will replace the current repo.

  `
	UserHomeDir, _ = os.UserHomeDir()
	UserTsdmDir = UserHomeDir + string(os.PathSeparator) + ".tsdm"
	cli := raspberry.Cli{
		AcceptedCommands: []string{"-v", "version", "-h", "help", "setup", "get"},
		HelpMsg:          help,
		Version:          0.1,
	}
	cli.Setup()
	cli.SetHandler("setup", setupDirs)
	cli.SetHandler("get", getRepo)
}
