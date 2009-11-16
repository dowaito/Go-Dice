package main

import (
	"time";
	"rand";
	"os";
	"strings";
	"strconv";
)

func roll(dtype, times int) int {
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

func rollTest(dtype, times, testBatch int) string {
	var results string = "";
	rolls := make([]string, times);
	for x := 0; x < testBatch; x++ {
		var total int = 0;
		var result string = "[";
		for i := 0; i < times; i++ {
			var thisRoll int;
			thisRoll = roll(dtype,1);
			rolls[i] = strconv.Itoa(thisRoll);
			total += thisRoll;
		}
		result += strings.Join(rolls, ", ");
		result += "]: ";
		result += strconv.Itoa(total);
		result += "\n";
		results += result;
	}
	return results;
}

func allTest() string {
	var output string = "Dice tests:\n";
	output += "3d6: ";
	output += rollTest(6,3,1);
	output += "1d20: ";
	output += rollTest(20,1,1);
	output += "2d10: ";
	output += rollTest(10,2,1);
	output += "3dF: ";
	output += rollTest(0,3,1);
	return output;
}

func main() {
	rand.Seed(time.Nanoseconds());
	os.Stdout.WriteString(allTest());
}