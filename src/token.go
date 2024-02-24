package main

import (
	"fmt"
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
	word_void_arg_empty     = iota
	//word_check_mod          = iota

	word_op_return = iota
	word_op_print  = iota
	word_op_comment = iota
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
	//case word_check_mod:
	//	return "word_check_mod"
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
	case word_op_comment:
		return "comment"
	default:
		//TODO ERROR
		fmt.Sprintf("Unknown mod value: %d", value)
		return ""
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
		//case "==":
		//	temp[i].typeOfToken = word_check_mod
		//	break
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
		case "//":
			temp[i].typeOfToken = word_op_comment
		default:
			//TODO error
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
