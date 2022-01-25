package main

import (
	"github.com/xBalbinus/polygonburn-watcher/daemon/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
