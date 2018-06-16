package files

import (
  "os"
)

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

func Delete(path string) (b bool, e error) {
  err := os.Remove(path)
  if err != nil {
    return false, err
  } else {
    return true, nil
  }
}
