package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type mod int

const (
	op_add  mod = iota
	op_sub  mod = iota
	op_div  mod = iota
	op_mul  mod = iota
	op_pow  mod = iota
	op_mod  mod = iota
	op_root mod = iota

	mod_par_cur_op mod = iota
	mod_par_cur_cl mod = iota
	mod_par_nor_op mod = iota
	mod_par_nor_cl mod = iota

	word_number mod = iota
	word_word   mod = iota

	word_assigne        mod = iota
	word_check_mod          = iota
	word_void_arg_empty     = iota

	word_op_return = iota
	word_op_print  = iota
	//word_op_if     = iota
	//word_op_else   = iota

	EOL = iota
	EOF = iota
)

type token struct {
	content     string
	typeOfToken mod
	//children    []token
}

func splitIntoTokens(fileInput string) (error, []token) {
	tokenList := []token{}

	for i := 0; i < len(strings.Split(fileInput, "\n")); i++ {
		for lineIter := 0; lineIter < len(strings.Split(fileInput, "\n")[i]); lineIter++ {
			var tokenContent string = ""
			var tokenType mod = 0
			var currentChar = strings.Split(fileInput, "\n")[i][lineIter]
			for string(currentChar) != " " {
				tokenContent += string(currentChar)

				if unicode.IsDigit(rune(currentChar)) {
					tokenType = word_number
				} else {
					tokenType = word_word
					if lineIter+1 < len(strings.Split(fileInput, "\n")[i])-1 {
						if currentChar == '-' && unicode.IsDigit(rune(strings.Split(fileInput, "\n")[i][lineIter+1])) {
							tokenType = word_number
						}
					}
				}
				lineIter++
				if lineIter < len(strings.Split(fileInput, "\n")[i]) {
					currentChar = strings.Split(fileInput, "\n")[i][lineIter]
				} else {
					break
				}
			}
			currentToken := token{
				content:     tokenContent,
				typeOfToken: tokenType,
			}
			tokenList = append(tokenList, currentToken)
		}
		tokenList = append(tokenList, token{content: "\n", typeOfToken: EOL})
	}
	tokenList = append(tokenList, token{content: "\n", typeOfToken: EOF})

	return nil, tokenList
}

func cleanFile(fileInput string) (error, string) {
	//check for \t and multi " " --> remove
	temp := strings.Replace(fileInput, "\t", " ", -1)
	temp = strings.Replace(fileInput, "    ", "", -1)
	return nil, temp
}

func enumToString(value mod) string {
	switch value {
	case op_add:
		return "op_add"
	case op_sub:
		return "op_sub"
	case op_div:
		return "op_div"
	case op_mul:
		return "op_mul"
	case mod_par_cur_op:
		return "mod_par_cur_op"
	case mod_par_cur_cl:
		return "mod_par_cur_cl"
	case mod_par_nor_op:
		return "mod_par_nor_op"
	case mod_par_nor_cl:
		return "mod_par_nor_cl"
	case word_number:
		return "word_number"
	case word_word:
		return "word_word"
	case word_assigne:
		return "word_assigne"
	case word_check_mod:
		return "word_check_mod"
	case word_void_arg_empty:
		return "word_void_arg_empty"
	case word_op_return:
		return "word_op_return"
	case word_op_print:
		return "word_op_print"
	case EOL:
		return "EOL"
	case EOF:
		return "EOF"
	case op_mod:
		return "%"
	case op_root:
		return "root"
	case op_pow:
		return "pow"
	default:
		return fmt.Sprintf("Unknown mod value: %d", value)
	}
}

func correctTypes(tokens []token) (error, []token) {
	temp := tokens
	for i := 0; i < len(temp); i++ {
		switch temp[i].content {
		case "=":
			temp[i].typeOfToken = word_assigne
			break
		case "{":
			temp[i].typeOfToken = mod_par_cur_op
			break
		case "}":
			temp[i].typeOfToken = mod_par_cur_cl
			break
		case "()":
			temp[i].typeOfToken = word_void_arg_empty
			break
		case "(":
			temp[i].typeOfToken = mod_par_nor_op
			break
		case ")":
			temp[i].typeOfToken = mod_par_nor_cl
			break
		case "==":
			temp[i].typeOfToken = word_check_mod
			break
		case "*":
			temp[i].typeOfToken = op_mul
			break
		case "/":
			temp[i].typeOfToken = op_div
			break
		case "-":
			temp[i].typeOfToken = op_sub
			break
		case "+":
			temp[i].typeOfToken = op_add
			break
		case "return":
			temp[i].typeOfToken = word_op_return
			break
		case "print":
			temp[i].typeOfToken = word_op_print
			break
		case "mod":
			temp[i].typeOfToken = op_mod
			break
		case "root":
			temp[i].typeOfToken = op_root
			break
		case "pow":
			temp[i].typeOfToken = op_pow
			break
		default:

			break
		}
	}
	return nil, temp
}

func tokenPrint(tokens []token) {
	for x := 0; x < len(tokens); x++ {
		tok := enumToString(tokens[x].typeOfToken)

		wor := tokens[x].content
		if wor == "\n" {
			wor = "\\n"
		}
		out := "token: " + wor + "\t ; with type: " + tok
		fmt.Println(out)
	}
	fmt.Println()
}

func main() {
	body, err := ioutil.ReadFile("./main.math" /*os.Args[1]*/)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	//get File

	err, out := cleanFile(string(body))
	//clear file of things

	if err != nil {
		log.Fatalf("cant clean file: %v", err)
	}
	//error if can't clean File

	body = []byte(out)
	//create body

	err, tokens := splitIntoTokens(string(body))
	//convert

	if err != nil {
		log.Fatalf("cant split into tokens: %v", err)
	}
	//throw error if can't convert
	err, correctedTokens := correctTypes(tokens)

	if err != nil {
		log.Fatalf("cant convert types: %v", err)
	}

	var newList []token
	for i := 0; i < len(correctedTokens) && correctedTokens[i].typeOfToken != EOL; i++ {
		newList = append(newList, correctedTokens[i])
	} //get first line

	tokens_, _, _ := parentheseLine(newList)
	tokenPrint(tokens_)
	//solve parentheses
}

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

type tokenTree struct { //idk
	content     string
	typeOfToken mod
	children    []tokenTree
}

/*
- read file
- remove white space
- get tokens:
lexer:
	- split into tokens
	- put into syntax-tree
	- index functions
- check if tokens make sens
- translate to logic
- make programm --> asm??
*/
/*
ops to add:
- abs ...
- power ... ...
- root ... ...
- factor ...
- mod ... ...
*/
