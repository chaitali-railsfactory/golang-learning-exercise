// If you want to build something using a Raspberry Pi, you'll probably use resistors. For this exercise, you need to know two things about them:

// Each resistor has a resistance value.
// Resistors are small - so small in fact that if you printed the resistance value on them, it would be hard to read.
// To get around this problem, manufacturers print color-coded bands onto the resistors to denote their resistance values. Each band has a position and a numeric value.

// The first 2 bands of a resistor have a simple encoding scheme: each color maps to a single number. For example, if they printed a brown band (value 1) followed by a green band (value 5), it would translate to the number 15.

// In this exercise you are going to create a helpful program so that you don't have to remember the values of the bands. The program will take color names as input and output a two digit number, even if the input is more than two colors!

// The band colors are encoded as follows:

// Black: 0
// Brown: 1
// Red: 2
// Orange: 3
// Yellow: 4
// Green: 5
// Blue: 6
// Violet: 7
// Grey: 8
// White: 9
// From the example above: brown-green should return 15 brown-green-violet should return 15 too, ignoring the third color.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var colorsMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Value(colors []string) int {
	var colorCombo string
	for _, color := range colors {
		color = strings.ToLower(color)
		colorCombo += strconv.Itoa(colorsMap[color])
	}
	value, err := strconv.Atoi(colorCombo)
	if err != nil {
		panic(err)
	}
	if value > 99 {
		value = value / 10
	}
	return value
}

var testCases = []struct {
	description string
	input       []string
	expected    int
}{

	{
		description: "Brown and black",
		input:       []string{"brown", "black"},
		expected:    10,
	},
	{
		description: "Blue and grey",
		input:       []string{"blue", "grey"},
		expected:    68,
	},
	{
		description: "Yellow and violet",
		input:       []string{"yellow", "violet"},
		expected:    47,
	},
	{
		description: "White and red",
		input:       []string{"white", "red"},
		expected:    92,
	},
	{
		description: "Orange and orange",
		input:       []string{"orange", "orange"},
		expected:    33,
	},
	{
		description: "Ignore additional colors",
		input:       []string{"green", "brown", "orange"},
		expected:    51,
	},
	{
		description: "Black and brown, one-digit",
		input:       []string{"black", "brown"},
		expected:    1,
	},
}

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	for _, testCase := range testCases {
		fmt.Printf("----------- \n")
		fmt.Printf("Description: %s \n", testCase.description)

		got := Value(testCase.input)

		if testCase.expected != got {
			fmt.Printf("%v Test Case Fail - Expected: %v and Got: %v \n", string(colorRed), testCase.expected, got)
		} else {
			fmt.Printf("%v Test Case Pass - Expected: %v and Got: %v \n", string(colorGreen), testCase.expected, got)
		}

	}
}
