package folder

import (
  "os"
  "github.com/njfix6/tunnel/pkg/file"
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

func Exists(path string) bool {
  return file.Exists(path)
}

func Delete(path string) (b bool, e error) {
  err := os.Remove(path)
  if err != nil {
    return false, err
  } else {
    return true, nil
  }
}

func Size(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}
