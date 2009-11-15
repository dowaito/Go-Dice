all: dice

dice: dice.8
	8l -o dice dice.8

dice.8:
	8g dice.go

clean:
	rm dice
	rm dice.8