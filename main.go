package main

import (
  "github.com/pandademic/raspberry"
)

func main(){
  help := `
           TSDM HELP 
         --------------

         the simple dotfile manager

         AVAILABLE COMMANDS:

         - help

         - -h

         - version

         - -v 


  `
  cli := raspberry.Cli{
    AcceptedCommands: []string{"-v","version","-h","help"},
    HelpMsg:help,
    Version: 0.1,
  }
  cli.Setup()
}
