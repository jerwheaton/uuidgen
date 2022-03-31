package main

import (
	"github.com/jerwheaton/uuidgen/cmd"
)

func main() {
	c := cmd.RunCommand()
	if err := c.Execute(); err != nil {
		panic(err)
	}
}
