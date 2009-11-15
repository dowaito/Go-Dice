package main

import (
       "fmt";
	"time";
	"rand";
)

func d6(x int) int {
	roll := 0;
	for i := 0; i < x; i++ {
		roll = roll + rand.Intn(5) + 1;
	}
	return roll;
}

func main() {
	for i := 0; i < 26; i++ {
		rand.Seed(time.Nanoseconds());
		fmt.Println(d6(3));
	}
}