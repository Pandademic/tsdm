package internal

import(
  "fmt"
  "os"
)

func Fatal(e error) {
  fmt.Println("Fatal:",msg)
  os.Exit(1)
}
