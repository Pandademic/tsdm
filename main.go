package main

import (
	"fmt"

	"github.com/Pandademic/luve"
	"github.com/go-git/go-git/v5"
	"github.com/pandademic/raspberry"
	"github.com/spf13/viper"

	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	UserHomeDir string
	UserTsdmDir string
	version     float64
)

func setupDirs() {
	fmt.Println("! 'setup' is depreacated and will be removed in v0.3 !")
	os.Exit(1)
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
func updateDots() {
	repoDir := UserTsdmDir + string(os.PathSeparator) + "dotfile-repo"
	os.Chdir(repoDir)
	r, err := git.PlainOpen(repoDir)
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}
	w, err := r.Worktree()
	if err != nil {
		fmt.Println("Fatal:", err)
		os.Exit(1)
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		fmt.Println("Fatal:", err)
		os.Exit(1)
	} else {
		fmt.Println("updated repo sucessfully!")
		os.Exit(0)
	}
}

type mytype struct {
	Name     string `maspstructure:"name"`
	Location string `mapstructure:"location"`
}

func syncFilesInternal(path string, file os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	o_s := runtime.GOOS
	filesParsed := make(map[string]mytype)
	err = viper.UnmarshalKey(o_s+".files", &filesParsed)
	if err != nil {
		panic(err)
	}
	fileName := filesParsed[path].Name
	fileLoc := filesParsed[path].Location
	if fileName != "" || fileLoc != "" {
		fmt.Println("Copying", fileName, "to", fileLoc)
		fileLoc = strings.Replace(fileLoc, "~", UserHomeDir, -1)
		original, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer original.Close()

		new, err := os.Create(fileLoc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer new.Close()
		b, err := io.Copy(new, original)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		luve.Luve(b)
		fmt.Println("copied!")
	}
	return nil
}
func syncFiles() {
	repoDir := UserTsdmDir + string(os.PathSeparator) + "dotfile-repo"
	os.Chdir(repoDir)

	viper.SetConfigName("tsdmrc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal:", err)
		os.Exit(1)
	}
	requiredVersion := viper.GetFloat64("reqVer")
	if requiredVersion != version {
		fmt.Println("Fatal: These dotfiles are meant for tsdm", requiredVersion, "But you are using", version)
		os.Exit(1)
	}

	if err := filepath.Walk(".", syncFilesInternal); err != nil {
		fmt.Println("Fatal: ", err)
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

	- sync <- copy the files in '~/.tsdm/dotfile-repo' to the correct locations
  `
	version = 0.2
	UserHomeDir, _ = os.UserHomeDir()
	UserTsdmDir = UserHomeDir + string(os.PathSeparator) + ".tsdm"
	cli := raspberry.Cli{
		AcceptedCommands: []string{"-v", "version", "-h", "help", "setup", "get", "update", "sync"},
		HelpMsg:          help,
		Version:          version,
	}
	cli.Setup()
	cli.SetHandler("setup", setupDirs)
	cli.SetHandler("get", getRepo)
	cli.SetHandler("update", updateDots)
	cli.SetHandler("sync", syncFiles)
}
