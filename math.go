package main

import (
	"fmt"
	"math"
	"strconv"
)

func isOperation(tokenTypeInput mod) bool {
	if tokenTypeInput == op_add || tokenTypeInput == op_sub || tokenTypeInput == op_div || tokenTypeInput == op_mul {
		return true
	}
	return false
}

func testParentheseLine(tokenIn []token) ([]token, int) {

	var tokenSample []token

	var endIndex int = 0

	for i := 0; i < len(tokenIn); i++ {
		if tokenIn[i].typeOfToken == mod_par_nor_op {
			if i+1 < len(tokenIn) {
				temp, indx := testParentheseLine(tokenIn[i+1:])
				for p := 0; p < indx+1; p++ {
					tokenSample = append(tokenIn[:(i)], tokenIn[(i+1):]...)
				}
				tokenSample = insertSliceAt(tokenSample, temp, i)
			}
		} else if tokenIn[i].typeOfToken == mod_par_nor_cl {
			endIndex = i
			tokenSample = mathLine(tokenSample)
		} else {
			tokenSample = append(tokenSample, tokenIn[i])
		}
	}

	return tokenSample, (endIndex)
}

func parentheseLine(tokenIn []token /*, startIndex int, endIndex int*/) ([]token, int, int) {

	var tokenSample []token

	var currentStartIndex int = 0
	var currentEndIndex int = 0

	//run through list --> find first open parenthese --> find next closed one restart if new open found --> solve content --> do again

	for containsParenthese(tokenIn) {
		for i := 0; i < len(tokenIn); i++ {
			if tokenIn[i].typeOfToken == mod_par_nor_op {
				tokenSample = []token{}
				currentStartIndex = i
			} else if tokenIn[i].typeOfToken == mod_par_nor_cl {
				currentEndIndex = i
				temp := mathLine(tokenSample) //--> replace result in token
				tokenSample = []token{}
				for delIter := 0; delIter < (currentEndIndex-currentStartIndex)+1; delIter++ { //???
					tokenIn = append(tokenIn[:currentStartIndex], tokenIn[currentStartIndex+1:]...)
				}
				tokenIn = insertSliceAt(tokenIn, temp, currentStartIndex)
				currentEndIndex = 0
				currentStartIndex = 0
				//tokenPrint(tokenIn)
				break
			} else {
				tokenSample = append(tokenSample, tokenIn[i])
				currentEndIndex = i
			}
		}
	}

	/*what:

	- first char until closed parenthese = enclosed problem

	- get line x
	- find open parenthese x
		--> capture anything until new open (call self) or closed parenthese
		- when closed paenthese --> solve anythin inside
	- return solved
	- replaced solved:
		- return solved list, + where in original list (start and end index)
	*/

	return mathLine(tokenIn), currentStartIndex, currentEndIndex
}

func containsParenthese(tokenIn []token) bool {
	for i := 0; i < len(tokenIn); i++ {
		if tokenIn[i].typeOfToken == mod_par_nor_cl || tokenIn[i].typeOfToken == mod_par_nor_op {
			return true
		}
	}
	return false
}

func insertSliceAt(list, insert []token, index int) []token {
	result := make([]token, len(list)+len(insert))
	copy(result, list[:index])
	copy(result[index:], insert)
	copy(result[index+len(insert):], list[index:])
	return result
}

func mathLine(tokenIn []token) []token {
	if len(tokenIn) <= 1 {
		return tokenIn
	}
	/*TODO:
	- undestand math


	x resolve all multiplications
	x resolve all divisions

	x find len of search
	- resolve par:
		- recusriv--> call function again with slice of string
	*/

	/*
		look for first open parenthese
		--> create slice of string (with out opening parenthese)
		--> call function with slice
		--> look for first closing parenthese
		--> solve anything until parenthese
		--> return
	*/

	var tempList [3]token

	for x := 0; ; x++ {
		didSomeThing := false
		for i := 0; i < len(tokenIn)-2; i++ {
			tempList[0+0] = tokenIn[i+0]
			tempList[0+1] = tokenIn[i+1]
			tempList[0+2] = tokenIn[i+2]
			if tempList[0].typeOfToken == word_number && tempList[2].typeOfToken == word_number {
				val1, _ := strconv.ParseFloat(string(tempList[0].content), 64)
				val2, _ := strconv.ParseFloat(string(tempList[2].content), 64)
				var res float64
				if tempList[1].typeOfToken == op_pow {
					res = math.Pow(val1, val2)
				} else if tempList[1].typeOfToken == op_mod {
					res = math.Mod(val1, val2)
				} else if tempList[1].typeOfToken == op_root {
					res = math.Pow(val2, 1/val1)
					fmt.Println("hi")
				}
				if tempList[1].typeOfToken == op_pow || tempList[1].typeOfToken == op_root || tempList[1].typeOfToken == op_mod {
					tokenIn[i].content = strconv.FormatFloat(res, 'f', 4, 64)
					tokenIn[i].typeOfToken = word_number
					tokenIn = append(tokenIn[:i+1], tokenIn[i+1+1:]...)
					tokenIn = append(tokenIn[:i+1], tokenIn[i+1+1:]...)
					didSomeThing = true
					break
				}
			}
		}
		if !didSomeThing {
			break
		}
	}

	for x := 0; ; x++ {
		didSomeThing := false
		for i := 0; i < len(tokenIn)-2; i++ {
			tempList[0+0] = tokenIn[i+0]
			tempList[0+1] = tokenIn[i+1]
			tempList[0+2] = tokenIn[i+2]
			if tempList[0].typeOfToken == word_number && tempList[2].typeOfToken == word_number {
				val1, _ := strconv.ParseFloat(string(tempList[0].content), 64)
				val2, _ := strconv.ParseFloat(string(tempList[2].content), 64)
				var res float64
				if tempList[1].typeOfToken == op_mul {
					res = val1 * val2
				} else if tempList[1].typeOfToken == op_div {
					res = val1 / val2
				}
				if tempList[1].typeOfToken == op_mul || tempList[1].typeOfToken == op_div {
					tokenIn[i].content = strconv.FormatFloat(res, 'f', 4, 64)
					tokenIn[i].typeOfToken = word_number
					tokenIn = append(tokenIn[:i+1], tokenIn[i+1+1:]...)
					tokenIn = append(tokenIn[:i+1], tokenIn[i+1+1:]...)
					didSomeThing = true
					break
				}
			}
		}
		if !didSomeThing {
			break
		}
	}

	for x := 0; ; x++ {
		didSomeThing := false
		for i := 0; i < len(tokenIn)-2; i++ {
			tempList[0+0] = tokenIn[i+0]
			tempList[0+1] = tokenIn[i+1]
			tempList[0+2] = tokenIn[i+2]
			if tempList[0].typeOfToken == word_number && tempList[2].typeOfToken == word_number {
				val1, _ := strconv.ParseFloat(string(tempList[0].content), 64)
				val2, _ := strconv.ParseFloat(string(tempList[2].content), 64)
				var res float64
				if tempList[1].typeOfToken == op_add {
					res = val1 + val2
				} else if tempList[1].typeOfToken == op_sub {
					res = val1 - val2
				}
				if tempList[1].typeOfToken == op_add || tempList[1].typeOfToken == op_sub {
					tokenIn[i].content = strconv.FormatFloat(res, 'f', 4, 64)
					tokenIn[i].typeOfToken = word_number
					tokenIn = append(tokenIn[:i+1], tokenIn[i+1+1:]...)
					tokenIn = append(tokenIn[:i+1], tokenIn[i+1+1:]...)
					didSomeThing = true
					break
				}
			}
		}
		if !didSomeThing {
			break
		}
	}

	return tokenIn
}

func factorial(number float64) float64 {

	if number == 1 || number < 1 {
		return 1
	}

	factorialOfNumber := number * factorial(number-1)

	return factorialOfNumber
}
