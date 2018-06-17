package main

import (
    "fmt"
    "os"
    "github.com/njfix6/tunnel/pkg/tunnel"
)

func main() {
  args := os.Args
  submain(args)
}

func submain(args []string) {
  if len(args) != 2 {
      fmt.Println("Usage: tunnel <name>")
      os.Exit(1)
    }
  app := args[1]
  tunnel.Dig(app)
}
