package tunnel

import (
  "github.com/njfix6/tunnel/pkg/folder"
)

func Dig (app string) {
  folder.Create(app)
  pkgPath := app + "/pkg"
  internalPath := app + "/internal"
  cmdPath := app + "/cmd"
  folder.Create(pkgPath)
  folder.Create(internalPath)
  folder.Create(cmdPath)
}
