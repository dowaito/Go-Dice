package dice

import (
	"time";
	"rand";
	"strings";
	"strconv";
)

func Roll(dformat string) int {
	var dtype, times int;
	var parsedDice []string = strings.Split(dformat, "d", 0);
	if parsedDice[1] == "F" || parsedDice[0] == "f" {
		dtype = 0;
	} else {
		dtype,_ = strconv.Atoi(parsedDice[1]);
	}
	times,_ = strconv.Atoi(parsedDice[0]);

	return Nroll(dtype, times);
}

func Nroll(dtype, times int) int {
	total := 0;
	for i := 0; i < times; i++ {
		rand.Seed(time.Nanoseconds());
		if dtype == 0 {
			total += rollFudge(); 
		}else{
			total += rand.Intn(dtype) + 1;
		}
	}
	return total
}

func rollFudge() int {
	fudgeDie := [3]int{-1, 0, 1};
	return fudgeDie[rand.Intn(3)];
}


