# Tunnel

## Requirements
1. [Go](https://golang.org/)

## Install
Run: `go get github.com/njfix6/tunnel/cmd/tunnel` to install command line tunnel

## Command Line Usage
Run: `tunnel <name>` where name is the name of the project you want to start

## Using the package
Import using `import "github.com/njfix6/tunnel"` in your file.

## Examples
```
import "github.com/njfix6/tunnel"
func main(){
  tunnel.Dig("TestApp")
}
```
