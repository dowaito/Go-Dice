all: roll

roll: dice.8 roll.8
	8l -o roll dice.8 roll.8

roll.8: roll.go
	8g roll.go

dice.8: dice.go
	8g dice.go
clean:
	rm *.8 roll