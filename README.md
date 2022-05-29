# tsdm
The Simple Dotfile Manager. 

## Architecture

Tsdm manages your dotfiles without symlinks , like so:

```
You pull dotfiles from git(one set at a time)
|
|
|
They are stored in your .tsdm directory ----------> you sync them , when they are copied to the specifed locations
                                                    |
                                                    |
                                                    |
                                                    you run update , where they are git pulled ---------> you repeat

```

You configure how tsdm manages your dotfiles , with a `tsdmrc.yml` file in your dotfile repo.

Example:

```yaml
reqVer: 0.1 # what is the minimum tsdm version these dotfiles need
windows: # OS block. The following instructions will only be executed on this OS
  files:
     - pwsh.txt: # name of the file.This is case senstive , tsdm parses the names of the files in dotfile directory , and looks for the data in the rc file 
        name: "powershell note" # this is used for info prinintg
        location: ~/foo/bar # where to copy it to on sync. Yes , it supports tilda's
        commentary: "this is optional , TSDM doesn't use it , its good for readability"
# etc. OS blocks

```
