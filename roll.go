package main

import (
	"./dice";
	"fmt";
	"flag";
	"strings";
	"strconv";
)

var format = flag.Bool("f", false, "Print roll with list of individual rolls")

func main() {
	flag.Parse();
	var dtype, times int;
	dtype, times = parseRoll(strings.ToLower(flag.Arg(0)));

	if !*format {
		fmt.Print(dice.Roll(dtype, times), "\n");
	} else {
		var roll, total int;
		rolls := make([]string,times);
		for i := 0; i < times; i++ {
			roll = dice.Roll(dtype, 1);
			rolls[i] = strconv.Itoa(roll);
			total += roll;
		}
		fmt.Print("[",strings.Join(rolls,", "),"] ", total, "\n");
	}
}

func parseRoll(dformat string) (int,int) {
	var dtype, times int;
	var parsedDice []string = strings.Split(dformat, "d", 0);
	if parsedDice[1] == "F" || parsedDice[0] == "f" {
		dtype = 0;
	} else {
		dtype,_ = strconv.Atoi(parsedDice[1]);
	}
	times,_ = strconv.Atoi(parsedDice[0]);

	return dtype, times;
}

