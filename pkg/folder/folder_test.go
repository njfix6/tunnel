package folder

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMakeDelete(t *testing.T) {
  assert := assert.New(t)
  filename := "go-lang-folder-test"
  Delete(filename)
  created, _ := Create(filename)
  if created != true {
    t.Error("Folder was not created")
  }
  exists := Exists(filename)
  assert.Equal(exists, true)
  created, _ = Create(filename)
  if created != false {
    t.Error("Folder created again")
  }
  removed, _ := Delete(filename)
  if removed != true {
    t.Error("Folder wasn't removed")
  }
  removed, _ = Delete(filename)
  if removed != false {
    t.Error("Folder removed again")
  }
}
