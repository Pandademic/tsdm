package main

// TODO ADD ENV VARS TO PATH FOR SUCCESS
import (
	"fmt"
	"github.com/Pandademic/luve"
	"github.com/pandademic/raspberry"
	"os"
)

var (
	UserHomeDir string
)

func setupDirs() {
	doesDirExist, err := os.Stat(UserHomeDir + "/.tsdm")
	UserHomeDir, _ = os.UserHomeDir()
	luve.Luve(doesDirExist)
	if os.IsNotExist(err) {
		if err := os.Mkdir(UserHomeDir+"/.tsdm", os.ModePerm); err != nil {
			fmt.Println("error creating directorys")
			fmt.Println("ERROR:", err)
			os.Exit(1)
		} else {
			fmt.Println("Success!")
		}
	} else {
		fmt.Println("Directory already exists!")
		os.Exit(1)
	}
}
func main() {
	help := `
           TSDM HELP 
         --------------

         the simple dotfile manager

         AVAILABLE COMMANDS:

         - help

         - -h

         - version

         - -v 
        
         - setup

  `
	cli := raspberry.Cli{
		AcceptedCommands: []string{"-v", "version", "-h", "help", "setup"},
		HelpMsg:          help,
		Version:          0.1,
	}
	cli.Setup()
	cli.SetHandler("setup", setupDirs)
}
