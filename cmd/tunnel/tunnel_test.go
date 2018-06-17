package main

import (
  "testing"
  "github.com/njfix6/tunnel/pkg/folder"
)

func cleanUp(app string){
  folder.Delete(app+"/pkg")
  folder.Delete(app+"/internal")
  folder.Delete(app+"/cmd")
  folder.Delete(app)
}

func TestMCommand(t *testing.T) {
  app := "go-lang-app-test"
  cleanUp(app)
  args := []string{"test", app}
  submain(args)
  cleanUp(app)
}
