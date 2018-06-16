package folder

import (
  "testing"
)

func TestMakeDelete(t *testing.T) {
  filename := "go-lang-folder-test"
  Delete(filename)
  created, _ := Create(filename)
  if created != true {
    t.Error("Folder was not created")
  }
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
