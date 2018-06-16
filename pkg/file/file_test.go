package file

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMakeDelete(t *testing.T) {
  assert := assert.New(t)
  filename := "go-lang-file-test"
  Delete(filename)
  created, _ := Create(filename)
  if created != true {
    t.Error("File was not created")
  }
  exists := Exists(filename)
  assert.Equal(true, exists, "The file should exist")
  created, _ = Create(filename)
  if created != false {
    t.Error("File created again")
  }
  removed, _ := Delete(filename)
  if removed != true {
    t.Error("File wasn't removed")
  }
  removed, _ = Delete(filename)
  if removed != false {
    t.Error("File removed again")
  }
}
