package main

import "strconv"

func process(fileIn []token) {
	var lineOffset int = 0

	//use const values and variable values

	values := map[string]float64{
		"pi": 3.14159265358973236,
		"e":  2.71828182845904523,
		"r2": 1.41421356237309504,
		"gr": 1.61803398874989484,
	}
	for {
		var currentLine []token = []token{}
		for i := 0; fileIn[i+lineOffset].typeOfToken != EOL; i++ {
			currentLine = append(currentLine, fileIn[i])
		}

		if currentLine[0].typeOfToken == word_op_comment {
			continue
		}
		if len(currentLine) >= 3 {
			if currentLine[0].typeOfToken == word_word &&
				currentLine[1].typeOfToken == word_assigne &&
				currentLine[2].typeOfToken == word_number {
				values[currentLine[0].content], _ = strconv.ParseFloat(string(currentLine[2].content), 64)
			}
		}
	}
}
