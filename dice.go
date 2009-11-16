package dice

import (
	"time";
	"rand";
)

func Roll(dtype, times int) int {
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


