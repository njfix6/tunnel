package tunnel

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/njfix6/tunnel/pkg/folder"
)



func cleanUp(app string){
  folder.Delete(app+"/pkg")
  folder.Delete(app+"/internal")
  folder.Delete(app+"/cmd")
  folder.Delete(app)
}

func TestMakeDelete(t *testing.T) {
  assert := assert.New(t)
  app := "go-lang-app-test"

  cleanUp(app)

  Dig(app)
  assert.Equal(true, folder.Exists(app) , "app folder should exist")
  assert.Equal(true, folder.Exists(app+"/pkg") , "pkg folder should exist")
  assert.Equal(true, folder.Exists(app+"/internal"), "internal folder should exist")
  assert.Equal(true, folder.Exists(app+"/cmd"), "cmd folder should exist")

  cleanUp(app)
}
