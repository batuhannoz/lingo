/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/batuhannoz/lingo/cmd"
	"github.com/batuhannoz/lingo/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
