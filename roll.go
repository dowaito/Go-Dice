package main

import (
	"./dice";
	"fmt";
	"flag";
	"strings";
)


func main() {
	flag.Parse();
	fmt.Print(dice.Roll(strings.ToLower(flag.Arg(0))), "\n");
}