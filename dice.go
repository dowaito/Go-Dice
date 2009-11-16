/*
 The dice package implements dice rolling.
 Each "roll" consists of individual die rolls added togethor for the total.
*/
package dice

import (
	"time";
	"rand";
)

/*
 Roll takes a die type (dtype) and a number of times (times) to roll the die.
 Fudge dice are implemented as dype 0.
*/
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


