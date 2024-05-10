package main

type function struct {
	body        []token
	returnValue float64
	argument    float64
}

func process(fileIn []token) {
	//staticValues := map[string]float64{
	//	"pi": 3.14159265358973236,
	//	"e":  2.71828182845904523,
	//	"r2": 1.41421356237309504,
	//	"gr": 1.61803398874989484,
	//}

	//uservalues := getVariables(fileIn)
	//userFunctions := getFunctions(fileIn)

	/*
		iter file:
			- reasigne variables
			- run function
			- print
	*/

}

func runFunction(funcIn function, userValues map[string]float64, userFunctions []function, staticValues map[string]float64) float64 {
	return 0.0
}

func getFunctions(fileIn []token) []function {
	//check for pattern --> fill in rest
	return []function{}
}

func getVariables(fileIn []token) map[string]float64 {
	var lineOffset int = 0
	uservalues := map[string]float64{}
	for {
		var currentLine []token = []token{}
		for i := 0; fileIn[i+lineOffset].typeOfToken != EOL; i++ {
			currentLine = append(currentLine, fileIn[i])
		}

		if currentLine[0].typeOfToken == word_op_comment {
			continue
		}
		if len(currentLine) >= 3 {
			if currentLine[0].typeOfToken == word_word && currentLine[1].typeOfToken == word_assigne && currentLine[2].typeOfToken == word_number {
				//uservalues[currentLine[0].content], _ = strconv.ParseFloat(string(currentLine[2].content), 64)
				uservalues[currentLine[0].content] = float64(0.0)
			}
		}
	}
}
