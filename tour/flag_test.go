package main

import (
	"flag"
	"log"
	"testing"
)

func TestFlag(*testing.T) {
	var name string
	flag.StringVar(&name, "name", "", "帮助信息")
	flag.StringVar(&name, "n", "", "help")
	flag.Parse()
	log.Printf("name: %s", name)
}
