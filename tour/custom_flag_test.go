package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"testing"
)

type CustomerFlagType string

func (c *CustomerFlagType) String() string {
	return fmt.Sprint(*c)
}
func (c *CustomerFlagType) Set(value string) error {
	if len(*c) > 0 {
		return errors.New("name flag already set")
	}
	*c = CustomerFlagType("customer flag type to set->" + value)
	return nil
}

func TestCustomFlag(*testing.T) {
	var n CustomerFlagType
	flag.Var(&n, "c", "-c xxx")
	flag.Parse()
	log.Printf("customer args: %s", n)
}
