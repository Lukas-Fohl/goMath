package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type mod int

const (
	op_add mod = iota
	op_sub mod = iota
	op_div mod = iota
	op_mul mod = iota

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

	err, out := cleanFile(string(body))

	if err != nil {
		log.Fatalf("cant clean file: %v", err)
	}

	body = []byte(out)

	err, tokens := splitIntoTokens(string(body))

	if err != nil {
		log.Fatalf("cant split into tokens: %v", err)
	}
	err, correctedTokens := correctTypes(tokens)
	//tokenPrint(correctedTokens)

	var newList []token
	for i := 0; i < len(correctedTokens) && correctedTokens[i].typeOfToken != EOL; i++ {
		newList = append(newList, correctedTokens[i])
	} //slice

	tokenPrint(newList)

	tokenPrint(parentheseLine(newList))
}

func isOperation(tokenTypeInput mod) bool {
	if tokenTypeInput == op_add || tokenTypeInput == op_sub || tokenTypeInput == op_div || tokenTypeInput == op_mul {
		return true
	}
	return false
}

func parentheseLine(tokenIn []token) []token {

	var tokenSample []token

	for i := 0; i < len(tokenIn); i++ {
		if tokenIn[i].typeOfToken == mod_par_nor_op {
			if i+1 <= len(tokenIn)-1 {
				tokenPrint(tokenIn[i+1 : len(tokenIn)-1])
				parentheseLine(tokenIn[i+1 : len(tokenIn)-1])
				//put return into array
			}
		} else if tokenIn[i].typeOfToken == mod_par_nor_cl {
			tokenSample = mathLine(tokenSample)
		} else {
			tokenSample = append(tokenSample, tokenIn[i])
		}
	}

	return tokenSample
}

func mathLine(tokenIn []token) []token {
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

	//tokenPrint(tokenIn)

	return tokenIn
}

type tokenTree struct {
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
