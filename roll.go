package main

import (
	"./dice";
	"fmt";
	"flag";
)


func main() {
	flag.Parse();
	fmt.Print(dice.Roll(flag.Arg(0)), "\n");
}