package main

import (
	"flag"
	"log"
	"testing"
)

var name string

func TestSubFlag(*testing.T) {
	flag.Parse()

	goCli := flag.NewFlagSet("go", flag.ExitOnError)
	goCli.StringVar(&name, "g", "", "go run sub_flag -go")

	javaCli := flag.NewFlagSet("java", flag.ExitOnError)
	javaCli.StringVar(&name, "j", "", "go run sub_flag -java")

	argSli := flag.Args()
	if len(argSli) <= 1 {
		log.Printf("-h see usage")
		return
	}
	switch argSli[0] {
	case "go":
		_ = goCli.Parse(argSli[1:])

	case "java":
		_ = javaCli.Parse(argSli[1:])
	default:
		log.Printf("-h see usage")
		return
	}
	log.Printf("sub flag: %s", name)
}
