package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"fmt"
)

func cleanFile(fileInput string) (error, string) {
	//check for \t and multi " " --> remove
	temp := strings.Replace(fileInput, "\t", " ", -1)
	temp = strings.Replace(fileInput, "    ", "", -1)
	return nil, temp
}

func main() {
	body, err := ioutil.ReadFile(os.Args[1])
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

	values := map[string]float64{
		"pi":	3.14159265358973236,
		"e":	2.71828182845904523,
		"r2":	1.41421356237309504,
		"gr":	1.61803398874989484,
	}
	//--> var name in file as index + value as value
	//index file first

	fmt.Printf("%v\n",values["e"])

	//TEST for math
	var newList []token
	for i := 0; i < len(correctedTokens) && correctedTokens[i].typeOfToken != EOL; i++ {
		newList = append(newList, correctedTokens[i])
	} //get first line

	tokens_, _, _ := parentheseLine(newList)
	tokenPrint(tokens_)
	//solve parentheses
}
