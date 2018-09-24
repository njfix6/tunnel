package file

import (
  "os"
  "fmt"
  "io/ioutil"
)

func Read(path string) (string, error) {
  b, err := ioutil.ReadFile(path) // just pass the file name
  if err != nil {
      return "", err
  }
  str := string(b)
  return str, nil
}

func ReadBytes(path string) ([]byte, error) {
  jsonFile, err := os.Open(path)
  if err != nil {
		return nil, err
	}
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)
  return byteValue, nil
}

func Create(path string) (b bool, e error) {
  var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
      return false, err
    }
	  file.Close()
    return true, nil
	} else {
    return false, err
  }
}

func Exists(path string) (b bool) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Println("doesn't exist")
    return false
  } else if err == nil {
    return true
  }
  return false
}

func Delete(path string) (b bool, e error) {
  err := os.Remove(path)
  if err != nil {
    return false, err
  } else {
    return true, nil
  }
}
