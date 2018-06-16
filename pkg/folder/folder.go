package folder

import (
  "os"
)

func Create(path string) (b bool, e error) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err = os.MkdirAll(path, os.ModePerm)
    if err != nil {
      return false, err
    }
    return true, nil
  } else {
    return false, err
  }
}

func Delete(path string) (b bool, e error) {
  err := os.Remove(path)
  if err != nil {
    return false, err
  } else {
    return true, nil
  }
}
